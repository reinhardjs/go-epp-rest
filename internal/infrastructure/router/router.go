package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/controller"
)

func NewRouter(appController controller.AppController) *gin.Engine {
	router := gin.Default()
	router.GET("/domain/check", func(c *gin.Context) { appController.Domain.Check(c) })
	router.GET("/contact/check", func(c *gin.Context) { appController.Contact.Check(c) })
	router.GET("/host/check", func(c *gin.Context) { appController.Host.Check(c) })
	return router
}
