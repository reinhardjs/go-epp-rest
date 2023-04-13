package utils

import (
	"log"
	"net"
	"sync"
	"time"
)

const (
	HELLO_PERIOD_IN_MINUTES = 5
)

// Session is a wrapper for a single tcp connection
type Session struct {
	Id          string     // A unique id to identify a connection
	conn        net.Conn   // The underlying TCP connection
	mu          sync.Mutex // mutex to prevent race conditions
	renewLock   sync.Mutex // mutex to prevent race conditions
	updateLock  sync.Mutex
	updateCond  *sync.Cond
	Pool        *SessionPool
	onUpdate    bool
	shouldLogin bool
	helloPeriod time.Duration
}

func (t *Session) helloWorker(period time.Duration, quit chan bool) {
	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Sending Hello")
			response, err := t.Pool.eppClient.SendHello(t.GetTcpConn())
			if err != nil {
				log.Println("Error sending hello:", err)
			} else {
				log.Println("Hello response:")
				log.Println(string(response))
			}
		case <-quit:
			return
		}
	}
}

func (t *Session) RunHelloWorker() {
	period := time.Duration(HELLO_PERIOD_IN_MINUTES) * time.Minute
	quit := make(chan bool)

	go t.helloWorker(period, quit)
}

func (t *Session) SetOnUpdate(onUpdate bool) {
	t.updateLock.Lock()
	t.onUpdate = onUpdate
	t.updateLock.Unlock()
}

func (t *Session) SetConn(conn net.Conn) {
	t.updateLock.Lock()
	t.conn = conn
	t.onUpdate = false
	t.updateCond.Signal()
	t.updateLock.Unlock()
}

func (t *Session) GetTcpConn() net.Conn {
	var conn net.Conn

	t.updateLock.Lock()

	if t.onUpdate {
		t.updateCond.Wait()
	}

	conn = t.conn
	t.updateLock.Unlock()

	return conn
}
