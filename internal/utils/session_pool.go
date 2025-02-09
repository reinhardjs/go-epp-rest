package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/constants"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/error_types"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp"
)

const maxQueueLength = 10_000

// connRequest wraps a channel to receive a connection
// and a channel to receive an error
type connRequest struct {
	connChan   chan *Session
	errChan    chan error
	isTimedout chan bool
}

type connRenewal struct {
	session *Session
	tcpConn chan net.Conn
	errChan chan error
}

// TcpConfig is a set of configuration for a TCP connection pool
type TcpConfig struct {
	Host         string
	Port         int
	TLSCert      *tls.Certificate
	RootCACert   *x509.CertPool
	MaxIdleConns int
	MaxOpenConn  int
}

// SessionPool represents a pool of tcp connections
type SessionPool struct {
	host         string
	port         int
	eppClient    adapter.EppClient
	tlsCert      *tls.Certificate
	rootCaCert   *x509.CertPool
	mu           sync.Mutex          // mutex to prevent race conditions
	idleConns    map[string]*Session // holds the idle connections
	numOpen      int                 // counter that tracks open connections
	maxOpenCount int
	maxIdleCount int
	renewChan    chan *connRenewal
	requestChan  chan *connRequest // A queue of connection requests
	generator    IDGenerator
}

// CreateTcpConnPool() creates a connection pool
// and starts the worker that handles connection request
func CreateTcpConnPool(cfg *TcpConfig) (*SessionPool, error) {

	pool := &SessionPool{
		host:         cfg.Host,
		port:         cfg.Port,
		tlsCert:      cfg.TLSCert,
		rootCaCert:   cfg.RootCACert,
		idleConns:    make(map[string]*Session),
		renewChan:    make(chan *connRenewal, cfg.MaxIdleConns),
		requestChan:  make(chan *connRequest, maxQueueLength),
		maxOpenCount: cfg.MaxOpenConn,
		maxIdleCount: cfg.MaxIdleConns,
		generator:    NewGenerator(),
	}

	go pool.handleConnectionRequest()

	go pool.handleConnectionRenewal()

	return pool, nil
}

func (p *SessionPool) SetEppClient(eppClient adapter.EppClient) {
	p.eppClient = eppClient
}

// put() attempts to return a used connection back to the pool
// It closes the connection if it can't do so
func (p *SessionPool) Put(c *Session) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.maxIdleCount > 0 && p.maxIdleCount > len(p.idleConns) {
		p.idleConns[c.Id] = c // put into the pool
	} else {
		if c.conn != nil {
			c.conn.Close()
		}
		p.numOpen--
	}
}

func (p *SessionPool) RenewTcpConn(c *Session) (net.Conn, error) {
	req := &connRenewal{
		session: c,
		tcpConn: make(chan net.Conn, 1),
		errChan: make(chan error, 1),
	}

	defer func() {
		// close channels
		close(req.tcpConn)
		close(req.errChan)
	}()

	p.renewChan <- req

	select {
	case tcpConn := <-req.tcpConn:
		return tcpConn, nil
	case err := <-req.errChan:
		return nil, err
	}
}

func (p *SessionPool) Init() error {
	p.mu.Lock()
	idleCons := p.idleConns
	maxIdleCount := p.maxIdleCount
	p.mu.Unlock()

	for len(idleCons) < maxIdleCount {
		p.mu.Lock()
		p.numOpen++
		p.mu.Unlock()

		newSession, err := p.createNewSession()
		if err != nil {
			p.mu.Lock()
			p.numOpen--
			p.mu.Unlock()
			return errors.Wrap(err, "SessionPool Init: p.createNewSession")
		}

		p.Put(newSession)
	}

	return nil
}

// get() retrieves a TCP connection
func (p *SessionPool) Get() (*Session, error) {
	p.mu.Lock()

	// Case 1: Gets a free connection from the pool if any
	numIdle := len(p.idleConns)
	if numIdle > 0 {
		// Loop map to get one conn
		for _, c := range p.idleConns {
			// remove from pool
			delete(p.idleConns, c.Id)
			p.mu.Unlock()
			return c, nil
		}
	}

	// Case 2: Queue a connection request
	if p.maxOpenCount > 0 && p.numOpen >= p.maxOpenCount {
		// Create the request
		req := &connRequest{
			connChan:   make(chan *Session, 1),
			errChan:    make(chan error, 1),
			isTimedout: make(chan bool, 1),
		}

		defer func() {
			// close channels
			close(req.connChan)
			close(req.errChan)
			close(req.isTimedout)
		}()

		// Queue the request
		p.requestChan <- req

		p.mu.Unlock()

		secondsTime, err := strconv.Atoi(os.Getenv(constants.REQUEST_TIMEOUT))
		if err != nil {
			req.errChan <- errors.New("REQUEST_TIMEOUT env value is not a valid number")
		}

		var (
			connSuccess = false
			hasTimeout  = false
			connFail    = false

			timeoutChan = time.After(time.Duration(secondsTime) * time.Second)
		)

		for {
			if connSuccess || hasTimeout || connFail {
				break
			}

			// Waits for either
			// 1. Request fulfilled, or
			// 2. An error is returned
			select {
			case <-timeoutChan:
				hasTimeout = true
				err := error_types.RequestTimeOutError{Detail: "request has timed out"}
				req.isTimedout <- true
				return nil, &err
			case tcpConn := <-req.connChan:
				connSuccess = true
				return tcpConn, nil
			case err := <-req.errChan:
				connFail = true
				return nil, err
			}
		}
	}

	// Case 3: Open a new connection
	p.numOpen++
	p.mu.Unlock()

	newSession, err := p.createNewSession()
	if err != nil {
		p.mu.Lock()
		p.numOpen--
		p.mu.Unlock()
		return nil, errors.Wrap(err, "SessionPool Get: p.createNewSession")
	}

	return newSession, nil
}

// createNewSession() creates a new TCP connection at p.host and p.port
func (p *SessionPool) createNewSession() (*Session, error) {
	c, err := p.openNewTcpConnection()
	if err != nil {
		return nil, err
	}

	session := &Session{
		// Use unix time as id
		Id:          p.generator.GenerateSessionId(),
		conn:        c,
		Pool:        p,
		shouldLogin: true, // should login for the first time
	}

	session.updateCond = sync.NewCond(&session.updateLock)

	session.RunHelloWorker()

	return session, nil
}

func (p *SessionPool) openNewTcpConnection() (net.Conn, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{*p.tlsCert},
		RootCAs:            p.rootCaCert,
	}
	addr := fmt.Sprintf("%s:%d", p.host, p.port)

	c, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return nil, errors.Wrap(err, "SessionPool openNewTcpConnection: tls.Dial")
	}

	// Read the greeting.
	greeting, err := registry_epp.ReadMessage(c)
	if err != nil {
		return nil, errors.Wrap(err, "SessionPool openNewTcpConnection: epp.ReadMessage")
	}

	log.Println(string(greeting))

	response, err := p.eppClient.DoLogin(c)
	if err != nil {
		return nil, errors.Wrap(err, "SessionPool openNewTcpConnection: p.login")
	}

	log.Println(string(response))

	return c, nil
}

// handleConnectionRequest() listens to the request queue
// and attempts to fulfil any incoming requests
func (p *SessionPool) handleConnectionRequest() {
	for req := range p.requestChan {
		secondsTime, err := strconv.Atoi(os.Getenv(constants.CONNECTION_REQUEST_TIMEOUT))
		if err != nil {
			req.errChan <- errors.New("CONNECTION_REQUEST_TIMEOUT env value is not a valid number")
		}

		var (
			requestDone = false
			hasTimeout  = false
			requestFail = false

			timeoutChan = time.After(time.Duration(secondsTime) * time.Second)
		)

		for {
			if requestDone || hasTimeout || requestFail {
				break
			}
			select {
			// request timeout
			case <-timeoutChan:
				hasTimeout = true
				err := error_types.RequestTimeOutError{Detail: "connection request has timed out"}
				req.errChan <- &err
			case <-req.isTimedout:
				hasTimeout = true
			default:
				p.mu.Lock()

				// First, we try to get an idle conn.
				// If fail, we try to open a new conn.
				// If both does not work, we try again in the next loop until timeout.
				numIdle := len(p.idleConns)
				if numIdle > 0 {
					for _, c := range p.idleConns {
						delete(p.idleConns, c.Id)
						p.mu.Unlock()
						req.connChan <- c // give conn
						requestDone = true
						break
					}
				} else if p.maxOpenCount > 0 && p.numOpen < p.maxOpenCount {
					p.numOpen++
					p.mu.Unlock()

					c, err := p.createNewSession()
					if err != nil {
						p.mu.Lock()
						p.numOpen--
						p.mu.Unlock()
						requestFail = true
						req.errChan <- err // give error
					} else {
						requestDone = true
						req.connChan <- c // give conn
					}
				} else {
					p.mu.Unlock()
				}
			}
		}
	}
}

func (p *SessionPool) handleConnectionRenewal() {
	for req := range p.renewChan {
		go func(req *connRenewal) {
			req.session.SetOnUpdate(true)

			conn, err := req.session.Pool.openNewTcpConnection()

			if err != nil {
				err = errors.Wrap(err, "SessionPool Get: req.session.pool.openNewTcpConnection")
				req.errChan <- err
				req.session.SetConn(nil)
			} else {
				req.tcpConn <- conn
				req.session.SetConn(conn)
			}

		}(req)
	}
}
