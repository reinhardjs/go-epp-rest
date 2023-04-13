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
	connChan chan *Session
	errChan  chan error
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
	host            string
	port            int
	eppClient       adapter.EppClient
	tlsCert         *tls.Certificate
	rootCaCert      *x509.CertPool
	mu              sync.Mutex          // mutex to prevent race conditions
	idleConns       map[string]*Session // holds the idle connections
	renewConns      map[string]*Session // holds session that needs to be renewed
	numOpen         int                 // counter that tracks open connections
	maxOpenCount    int
	maxIdleCount    int
	renewChan       chan *connRenewal
	requestChan     chan *connRequest // A queue of connection requests
	requestChanPool *sync.Pool
}

// CreateTcpConnPool() creates a connection pool
// and starts the worker that handles connection request
func CreateTcpConnPool(cfg *TcpConfig) (*SessionPool, error) {
	reqChanPool := sync.Pool{
		New: func() interface{} {
			return &connRequest{
				connChan: make(chan *Session, 1),
				errChan:  make(chan error, 1),
			}
		},
	}

	pool := &SessionPool{
		host:            cfg.Host,
		port:            cfg.Port,
		tlsCert:         cfg.TLSCert,
		rootCaCert:      cfg.RootCACert,
		idleConns:       make(map[string]*Session),
		renewConns:      make(map[string]*Session),
		renewChan:       make(chan *connRenewal, maxQueueLength),
		requestChan:     make(chan *connRequest, maxQueueLength),
		requestChanPool: &reqChanPool,
		maxOpenCount:    cfg.MaxOpenConn,
		maxIdleCount:    cfg.MaxIdleConns,
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
		if c.Conn != nil {
			c.Conn.Close()
		}
		p.numOpen--
	}
}

func (p *SessionPool) Retry(c *Session) (net.Conn, error) {
	c.renewLock.Lock()
	defer c.renewLock.Unlock()

	p.renewConns[c.Id] = c

	req := &connRenewal{
		session: c,
		tcpConn: make(chan net.Conn, 1),
		errChan: make(chan error, 1),
	}

	p.renewChan <- req

	select {
	case tcpConn := <-req.tcpConn:
		return tcpConn, nil
	case err := <-req.errChan:
		return nil, err
	}
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
		req := p.requestChanPool.Get().(*connRequest)

		// Queue the request
		p.requestChan <- req

		p.mu.Unlock()

		// put back request chan to the pool, for being re-used
		defer p.requestChanPool.Put(req)

		// Waits for either
		// 1. Request fulfilled, or
		// 2. An error is returned
		select {
		case tcpConn := <-req.connChan:
			return tcpConn, nil
		case err := <-req.errChan:
			return nil, err
		}
	}

	// Case 3: Open a new connection
	p.numOpen++
	p.mu.Unlock()

	newTcpConn, err := p.createNewSession()
	if err != nil {
		p.mu.Lock()
		p.numOpen--
		p.mu.Unlock()
		return nil, errors.Wrap(err, "SessionPool Get: p.openNewTcpConnection")
	}

	return newTcpConn, nil
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

// createNewSession() creates a new TCP connection at p.host and p.port
func (p *SessionPool) createNewSession() (*Session, error) {
	c, err := p.openNewTcpConnection()
	if err != nil {
		return nil, err
	}

	session := &Session{
		// Use unix time as id
		Id:          fmt.Sprintf("%v", time.Now().UnixNano()),
		Conn:        c,
		Pool:        p,
		shouldLogin: true, // should login for the first time
	}

	session.updateCond = sync.NewCond(&session.updateLock)

	return session, nil
}

// handleConnectionRequest() listens to the request queue
// and attempts to fulfil any incoming requests
func (p *SessionPool) handleConnectionRequest() {
	for req := range p.requestChan {
		go func(req *connRequest) {
			secondsTime, err := strconv.Atoi(os.Getenv(constants.REQUEST_TIMEOUT))
			if err != nil {
				req.errChan <- errors.New("REQUEST_TIMEOUT env value is not a valid number")
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
					err := error_types.RequestTimeOutError{Detail: "queue waiting time has timed out"}
					req.errChan <- &err
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
		}(req)
	}
}

func (p *SessionPool) handleConnectionRenewal() {
	for req := range p.renewChan {
		go func(req *connRenewal) {
			req.session.updateLock.Lock()
			req.session.onUpdate = true
			req.session.updateLock.Unlock()

			conn, err := req.session.Pool.openNewTcpConnection()

			if err != nil {
				err = errors.Wrap(err, "SessionPool Get: req.session.pool.openNewTcpConnection")
				req.errChan <- err
				req.session.RenewConn(nil)
			} else {
				req.tcpConn <- conn
				req.session.RenewConn(conn)
			}
		}(req)
	}
}
