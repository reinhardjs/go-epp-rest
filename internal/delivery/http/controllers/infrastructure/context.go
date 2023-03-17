package infrastructure

import "github.com/gin-gonic/gin"

type Context interface {
	// Close()
	BindQuery(obj any) error
	Query(key string) string
	AbortWithError(code int, fatalErr error) *gin.Error
}
