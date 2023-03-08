package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure/wrapper"
)

func NewRouter(appController controllers.AppController) *gin.Engine {
	router := gin.Default()

	router.GET("/domain/check", func(c *gin.Context) { appController.Domain.Check(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/domain/create", func(c *gin.Context) { appController.Domain.Create(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/domain/delete", func(c *gin.Context) { appController.Domain.Delete(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/domain/info", func(c *gin.Context) { appController.Domain.Info(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/domain/secdnsupdate", func(c *gin.Context) { appController.Domain.SecDNSUpdate(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/domain/contact/update", func(c *gin.Context) { appController.Domain.ContactUpdate(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/domain/status/update", func(c *gin.Context) { appController.Domain.StatusUpdate(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/domain/authinfo/update", func(c *gin.Context) { appController.Domain.AuthInfoUpdate(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/domain/nameserver/update", func(c *gin.Context) { appController.Domain.NameserverUpdate(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/domain/renew", func(c *gin.Context) { appController.Domain.Renew(&wrapper.ContextAdapter{Context: c}) })

	router.GET("/host/check", func(c *gin.Context) { appController.Host.Check(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/host/create", func(c *gin.Context) { appController.Host.Create(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/host/update", func(c *gin.Context) { appController.Host.Update(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/host/delete", func(c *gin.Context) { appController.Host.Delete(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/host/info", func(c *gin.Context) { appController.Host.Info(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/host/change", func(c *gin.Context) { appController.Host.Change(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/host/checkcreate", func(c *gin.Context) { appController.Host.CheckAndCreate(&wrapper.ContextAdapter{Context: c}) })

	router.GET("/contact/check", func(c *gin.Context) { appController.Contact.Check(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/contact/create", func(c *gin.Context) { appController.Contact.Create(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/contact/update", func(c *gin.Context) { appController.Contact.Update(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/contact/delete", func(c *gin.Context) { appController.Contact.Delete(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/contact/info", func(c *gin.Context) { appController.Contact.Info(&wrapper.ContextAdapter{Context: c}) })

	router.GET("/transfer/check", func(c *gin.Context) { appController.Transfer.Check(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/transfer/request", func(c *gin.Context) { appController.Transfer.Request(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/transfer/cancel", func(c *gin.Context) { appController.Transfer.Cancel(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/transfer/approve", func(c *gin.Context) { appController.Transfer.Approve(&wrapper.ContextAdapter{Context: c}) })
	router.GET("/transfer/reject", func(c *gin.Context) { appController.Transfer.Approve(&wrapper.ContextAdapter{Context: c}) })

	router.GET("/poll", func(c *gin.Context) { appController.Poll.Poll(&wrapper.ContextAdapter{Context: c}) })

	return router
}
