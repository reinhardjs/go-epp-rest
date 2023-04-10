package middlewares

import (
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func ServerErrorHandler(c *gin.Context, err any) {
	defer func() {
		runtime.GC()
	}()

	serverError := errors.Wrap(err.(error), "internal server error")
	log.Println(serverError)
	c.Abort()
	c.String(500, "2400 Command failed; Internal server error")
}
