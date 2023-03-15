package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"
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
func CreateTcpConnPool(cfg *TcpConfig) (*TcpConnPool, error) {
	pool := &TcpConnPool{
		host:         cfg.Host,
		port:         cfg.Port,
		tlsCert:      cfg.TLSCert,
		rootCaCert:   cfg.RootCACert,
		idleConns:    make(map[string]*TcpConn),
		requestChan:  make(chan *connRequest, maxQueueLength),
		maxOpenCount: cfg.MaxOpenConn,
		maxIdleCount: cfg.MaxIdleConns,
	}

	go pool.handleConnectionRequest()

	return pool, nil
}

// TcpConnPool represents a pool of tcp connections
type TcpConnPool struct {
	host         string
	port         int
	tlsCert      *tls.Certificate
	rootCaCert   *x509.CertPool
	mu           sync.Mutex          // mutex to prevent race conditions
	idleConns    map[string]*TcpConn // holds the idle connections
	numOpen      int                 // counter that tracks open connections
	maxOpenCount int
	maxIdleCount int
	requestChan  chan *connRequest // A queue of connection requests
}

// TcpConn is a wrapper for a single tcp connection
type TcpConn struct {
	Id   string       // A unique id to identify a connection
	Pool *TcpConnPool // The TCP connecion pool
	Conn net.Conn     // The underlying TCP connection
}

// connRequest wraps a channel to receive a connection
// and a channel to receive an error
type connRequest struct {
	connChan chan *TcpConn
	errChan  chan error
}

// put() attempts to return a used connection back to the pool
// It closes the connection if it can't do so
func (p *TcpConnPool) Put(c *TcpConn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.maxIdleCount > 0 && p.maxIdleCount > len(p.idleConns) {
		p.idleConns[c.Id] = c // put into the pool
	} else {
		c.Conn.Close()
		c.Pool.numOpen--
	}
}

// get() retrieves a TCP connection
func (p *TcpConnPool) Get() (*TcpConn, error) {
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
			connChan: make(chan *TcpConn, 1),
			errChan:  make(chan error, 1),
		}

		// Queue the request
		p.requestChan <- req

		p.mu.Unlock()

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
func (p *TcpConnPool) openNewTcpConnection() (*TcpConn, error) {
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
		Id:   fmt.Sprintf("%v", time.Now().UnixNano()),
		Conn: c,
		Pool: p,
	}, nil
}

// handleConnectionRequest() listens to the request queue
// and attempts to fulfil any incoming requests
func (p *TcpConnPool) handleConnectionRequest() {

	for req := range p.requestChan {
		log.Println("REQUEST CHAN", req)
		var (
			requestDone = false
			hasTimeout  = false

			// start a 3-second timeout
			timeoutChan = time.After(200 * time.Millisecond)
		)

		for {
			if requestDone || hasTimeout {
				break
			}
			select {
			// request timeout
			case <-timeoutChan:
				hasTimeout = true
				req.errChan <- errors.New("connection request timeout")
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
					} else {
						req.connChan <- c // give conn
						requestDone = true
					}
				} else {
					p.mu.Unlock()
				}
			}
		}
	}
}
