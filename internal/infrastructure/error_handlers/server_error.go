package error_handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func ServerErrorHandler(c *gin.Context, err any) {
	serverError := errors.Wrap(err.(error), "internal server error: ")
	log.Println(serverError)
	c.String(500, "2400\tCommand failed; Internal server error")
	c.Abort()
}
