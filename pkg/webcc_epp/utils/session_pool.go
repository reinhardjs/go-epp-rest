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
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp"
)

const maxQueueLength = 10_000

// TcpConfig is a set of configuration for a TCP connection pool
type TcpConfig struct {
	Host         string
	Port         int
	TLSCert      *tls.Certificate
	RootCACert   *x509.CertPool
	MaxIdleConns int
	MaxOpenConn  int
}

// CreateTcpConnPool() creates a connection pool
// and starts the worker that handles connection request
func CreateTcpConnPool(cfg *TcpConfig) (*SessionPool, error) {
	reqChanPool := sync.Pool{
		New: func() interface{} {
			return &connRequest{
				connChan: make(chan *TcpConn, 1),
				errChan:  make(chan error, 1),
			}
		},
	}

	pool := &SessionPool{
		host:            cfg.Host,
		port:            cfg.Port,
		tlsCert:         cfg.TLSCert,
		rootCaCert:      cfg.RootCACert,
		idleConns:       make(map[string]*TcpConn),
		requestChan:     make(chan *connRequest, maxQueueLength),
		requestChanPool: &reqChanPool,
		maxOpenCount:    cfg.MaxOpenConn,
		maxIdleCount:    cfg.MaxIdleConns,
	}

	go pool.handleConnectionRequest()

	return pool, nil
}

// SessionPool represents a pool of tcp connections
type SessionPool struct {
	host            string
	port            int
	tlsCert         *tls.Certificate
	rootCaCert      *x509.CertPool
	mu              sync.Mutex          // mutex to prevent race conditions
	idleConns       map[string]*TcpConn // holds the idle connections
	numOpen         int                 // counter that tracks open connections
	maxOpenCount    int
	maxIdleCount    int
	requestChan     chan *connRequest // A queue of connection requests
	requestChanPool *sync.Pool
}

// TcpConn is a wrapper for a single tcp connection
type TcpConn struct {
	Id          string     // A unique id to identify a connection
	Conn        net.Conn   // The underlying TCP connection
	mu          sync.Mutex // mutex to prevent race conditions
	shouldLogin bool
}

func (t *TcpConn) GetShouldLogin() bool {
	t.mu.Lock()
	shouldLogin := t.shouldLogin
	t.mu.Unlock()

	return shouldLogin
}

func (t *TcpConn) SetShouldLogin(shouldLogin bool) {
	t.mu.Lock()
	t.shouldLogin = shouldLogin
	t.mu.Unlock()
}

// connRequest wraps a channel to receive a connection
// and a channel to receive an error
type connRequest struct {
	connChan chan *TcpConn
	errChan  chan error
}

func (p *SessionPool) Throw() {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.numOpen--
}

// put() attempts to return a used connection back to the pool
// It closes the connection if it can't do so
func (p *SessionPool) Put(c *TcpConn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.maxIdleCount > 0 && p.maxIdleCount > len(p.idleConns) {
		p.idleConns[c.Id] = c // put into the pool
	} else {
		c.Conn.Close()
		p.numOpen--
	}
}

// get() retrieves a TCP connection
func (p *SessionPool) Get() (*TcpConn, error) {
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

	newTcpConn, err := p.openNewTcpConnection()
	if err != nil {
		p.mu.Lock()
		p.numOpen--
		p.mu.Unlock()
		return nil, errors.Wrap(err, "TcpConnPool Get: p.openNewTcpConnection")
	}

	return newTcpConn, nil
}

// openNewTcpConnection() creates a new TCP connection at p.host and p.port
func (p *SessionPool) openNewTcpConnection() (*TcpConn, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{*p.tlsCert},
		RootCAs:            p.rootCaCert,
	}
	addr := fmt.Sprintf("%s:%d", p.host, p.port)

	c, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return nil, errors.Wrap(err, "TcpConnPool openNewTcpConnection: tls.Dial")
	}

	// Read the greeting.
	greeting, err := registry_epp.ReadMessage(c)
	if err != nil {
		return nil, errors.Wrap(err, "TcpConnPool openNewTcpConnection: epp.ReadMessage")
	}

	log.Println(string(greeting))

	return &TcpConn{
		// Use unix time as id
		Id:          fmt.Sprintf("%v", time.Now().UnixNano()),
		Conn:        c,
		shouldLogin: true, // should login for the first time
	}, nil
}

// handleConnectionRequest() listens to the request queue
// and attempts to fulfil any incoming requests
func (p *SessionPool) handleConnectionRequest() {
	for req := range p.requestChan {
		go p.handle(req)
	}
}

func (p *SessionPool) handle(req *connRequest) {
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

				c, err := p.openNewTcpConnection()
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
