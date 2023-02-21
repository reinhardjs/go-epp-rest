package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/controller"
)

func NewRouter(appController controller.AppController) *gin.Engine {
	router := gin.Default()
	router.GET("/domain/check", func(c *gin.Context) { appController.Domain.CheckDomain(c) })
	return router
}
