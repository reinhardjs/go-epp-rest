package utils

import (
	"net"
	"sync"
)

// Session is a wrapper for a single tcp connection
type Session struct {
	Id          string     // A unique id to identify a connection
	Conn        net.Conn   // The underlying TCP connection
	mu          sync.Mutex // mutex to prevent race conditions
	renewLock   sync.Mutex // mutex to prevent race conditions
	updateLock  sync.Mutex
	updateCond  *sync.Cond
	Pool        *SessionPool
	onUpdate    bool
	shouldLogin bool
}

func (t *Session) RenewConn(conn net.Conn) {
	t.updateLock.Lock()
	t.Conn = conn
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

	conn = t.Conn
	t.updateLock.Unlock()

	return conn
}
