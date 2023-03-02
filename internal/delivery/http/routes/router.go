package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
)

func NewRouter(appController controllers.AppController) *gin.Engine {
	router := gin.Default()
	router.GET("/domain/check", func(c *gin.Context) { appController.Domain.Check(c) })
	router.GET("/contact/check", func(c *gin.Context) { appController.Contact.Check(c) })
	router.GET("/host/check", func(c *gin.Context) { appController.Host.Check(c) })

	router.GET("/host/create", func(c *gin.Context) { appController.Host.Create(c) })
	router.GET("/contact/create", func(c *gin.Context) { appController.Contact.Create(c) })
	router.GET("/domain/create", func(c *gin.Context) { appController.Domain.Create(c) })

	router.GET("/contact/update", func(c *gin.Context) { appController.Contact.Update(c) })
	router.GET("/host/update", func(c *gin.Context) { appController.Host.Update(c) })

	router.GET("/domain/delete", func(c *gin.Context) { appController.Domain.Delete(c) })
	router.GET("/contact/delete", func(c *gin.Context) { appController.Contact.Delete(c) })
	router.GET("/host/delete", func(c *gin.Context) { appController.Host.Delete(c) })

	router.GET("/domain/info", func(c *gin.Context) { appController.Domain.Info((c)) })
	router.GET("/host/info", func(c *gin.Context) { appController.Host.Info((c)) })
	router.GET("/contact/info", func(c *gin.Context) { appController.Contact.Info((c)) })

	router.GET("/transfer/check", func(c *gin.Context) { appController.Transfer.Check((c)) })
	router.GET("/transfer/request", func(c *gin.Context) { appController.Transfer.Request((c)) })
	router.GET("/transfer/cancel", func(c *gin.Context) { appController.Transfer.Cancel((c)) })
	router.GET("/transfer/approve", func(c *gin.Context) { appController.Transfer.Approve((c)) })
	router.GET("/transfer/reject", func(c *gin.Context) { appController.Transfer.Approve((c)) })

	router.GET("domain/secdnsupdate", func(c *gin.Context) { appController.Domain.SecDNSUpdate(c) })

	return router
}
