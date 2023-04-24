package utils

import (
	"fmt"
	"sync"
	"time"
)

type idGenerator struct {
	mu                    sync.Mutex
	requestDifferentiator int
	sessionDifferentiator int
}

type IDGenerator interface {
	GenerateRequestId() string
	GenerateSessionId() string
}

func NewGenerator() IDGenerator {
	return &idGenerator{
		requestDifferentiator: -1,
		sessionDifferentiator: -1,
	}
}

func (g *idGenerator) GenerateRequestId() string {
	g.mu.Lock()

	if g.requestDifferentiator == 9 {
		g.requestDifferentiator = 0
	} else {
		g.requestDifferentiator++
	}

	// create a time variable
	now := time.Now()

	// convert to unix time in nanoseconds
	unixNano := now.UnixNano()
	g.mu.Unlock()

	return fmt.Sprintf("request-%v-%v", unixNano, g.requestDifferentiator)
}

func (g *idGenerator) GenerateSessionId() string {
	g.mu.Lock()

	if g.sessionDifferentiator == 9 {
		g.sessionDifferentiator = 0
	} else {
		g.sessionDifferentiator++
	}

	// create a time variable
	now := time.Now()

	// convert to unix time in nanoseconds
	unixNano := now.UnixNano()
	g.mu.Unlock()

	return fmt.Sprintf("session-%v-%v", unixNano, g.sessionDifferentiator)
}
