package utils

import (
	"fmt"
	"sync"
	"time"
)

type generator struct {
	mu             sync.Mutex
	differentiator int
}

type Generator interface {
	GenerateRequestId() string
}

func NewGenerator() Generator {
	return &generator{
		differentiator: -1,
	}
}

func (g *generator) GenerateRequestId() string {
	g.mu.Lock()

	if g.differentiator == 9 {
		g.differentiator = 0
	} else {
		g.differentiator++
	}

	// create a time variable
	now := time.Now()

	// convert to unix time in nanoseconds
	unixNano := now.UnixNano()
	g.mu.Unlock()

	return fmt.Sprintf("request-%v-%v", unixNano, g.differentiator)
}
