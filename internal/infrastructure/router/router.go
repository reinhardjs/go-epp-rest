package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/controller"
)

func NewRouter(appController controller.AppController) *gin.Engine {
	router := gin.Default()
	return router
}
