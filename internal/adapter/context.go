package adapter

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

type ContextAdapter struct {
	*gin.Context
}

func (c *ContextAdapter) Close() {
	c.Request.Context().Done()
	c.Writer.Flush()
	runtime.GC()
}

func (c *ContextAdapter) AbortWithError(code int, fatalErr error) *gin.Error {
	return c.Context.AbortWithError(code, fatalErr)
}

func (c *ContextAdapter) BindQuery(obj any) error {
	return c.Context.BindQuery(obj)
}

func (c *ContextAdapter) Query(key string) string {
	return c.Context.Query(key)
}

func (c *ContextAdapter) String(code int, format string, values ...any) {
	c.Context.String(code, format)
}
