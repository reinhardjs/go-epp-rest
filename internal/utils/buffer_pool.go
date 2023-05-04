package utils

import (
	"bytes"
	"sync"
)

var bufferPoolLock = &sync.Mutex{}

var bufferPoolInstance *bufferPool

type bufferPool struct {
	pool sync.Pool
}

type BufferPool interface {
	Get() *bytes.Buffer
	Put(buffer *bytes.Buffer)
}

func GetBufferPoolInstance() BufferPool {
	bufferPoolLock.Lock()
	defer bufferPoolLock.Unlock()

	if bufferPoolInstance == nil {
		if bufferPoolInstance == nil {
			bufferPoolInstance = &bufferPool{
				pool: sync.Pool{
					New: func() interface{} {
						return &bytes.Buffer{}
					},
				},
			}

		}

		return bufferPoolInstance
	}

	return bufferPoolInstance
}

func (b *bufferPool) Get() *bytes.Buffer {
	return b.pool.Get().(*bytes.Buffer)
}

func (b *bufferPool) Put(buffer *bytes.Buffer) {
	buffer.Reset()
	b.pool.Put(buffer)
}
