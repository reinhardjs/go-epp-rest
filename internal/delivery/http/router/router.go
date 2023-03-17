package router

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"gitlab.com/merekmu/go-epp-rest/internal/adapter"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/middlewares"

	"github.com/gin-contrib/pprof"
)

type handler struct {
	appController controllers.AppController
}

func (h *handler) domainCheck(c *gin.Context) {
	h.appController.Domain.Check(&adapter.ContextAdapter{Context: c})
}

func (h *handler) domainCreate(c *gin.Context) {
	h.appController.Domain.Create(&adapter.ContextAdapter{Context: c})
}

func NewRouter(appController controllers.AppController) *gin.Engine {
	handler := &handler{appController}
	router := gin.Default()

	// Profiling monitor
	pprof.Register(router)

	// Prometheus monitoring
	prometheus := ginprometheus.NewPrometheus("gin")
	prometheus.Use(router)

	// Use timeout handler
	router.Use(middlewares.TimeoutMiddleware())

	// Use error handler middleware
	router.Use(middlewares.ClientErrorHandler)                     // Error related to client error
	router.Use(gin.CustomRecovery(middlewares.ServerErrorHandler)) // Error related to server error resulted from like panic/exception, etc..

	router.GET("/domain/check", handler.domainCheck)
	router.GET("/domain/create", handler.domainCreate)
	router.GET("/domain/delete", func(c *gin.Context) { appController.Domain.Delete(&adapter.ContextAdapter{Context: c}) })
	router.GET("/domain/info", func(c *gin.Context) { appController.Domain.Info(&adapter.ContextAdapter{Context: c}) })
	router.GET("/domain/secdnsupdate", func(c *gin.Context) { appController.Domain.SecDNSUpdate(&adapter.ContextAdapter{Context: c}) })
	router.GET("/domain/contact/update", func(c *gin.Context) { appController.Domain.ContactUpdate(&adapter.ContextAdapter{Context: c}) })
	router.GET("/domain/status/update", func(c *gin.Context) { appController.Domain.StatusUpdate(&adapter.ContextAdapter{Context: c}) })
	router.GET("/domain/authinfo/update", func(c *gin.Context) { appController.Domain.AuthInfoUpdate(&adapter.ContextAdapter{Context: c}) })
	router.GET("/domain/nameserver/update", func(c *gin.Context) { appController.Domain.NameserverUpdate(&adapter.ContextAdapter{Context: c}) })
	router.GET("/domain/renew", func(c *gin.Context) { appController.Domain.Renew(&adapter.ContextAdapter{Context: c}) })

	router.GET("/host/check", func(c *gin.Context) { appController.Host.Check(&adapter.ContextAdapter{Context: c}) })
	router.GET("/host/create", func(c *gin.Context) { appController.Host.Create(&adapter.ContextAdapter{Context: c}) })
	router.GET("/host/update", func(c *gin.Context) { appController.Host.Update(&adapter.ContextAdapter{Context: c}) })
	router.GET("/host/delete", func(c *gin.Context) { appController.Host.Delete(&adapter.ContextAdapter{Context: c}) })
	router.GET("/host/info", func(c *gin.Context) { appController.Host.Info(&adapter.ContextAdapter{Context: c}) })
	router.GET("/host/change", func(c *gin.Context) { appController.Host.Change(&adapter.ContextAdapter{Context: c}) })
	router.GET("/host/checkcreate", func(c *gin.Context) { appController.Host.CheckAndCreate(&adapter.ContextAdapter{Context: c}) })

	router.GET("/contact/check", func(c *gin.Context) { appController.Contact.Check(&adapter.ContextAdapter{Context: c}) })
	router.GET("/contact/create", func(c *gin.Context) { appController.Contact.Create(&adapter.ContextAdapter{Context: c}) })
	router.GET("/contact/update", func(c *gin.Context) { appController.Contact.Update(&adapter.ContextAdapter{Context: c}) })
	router.GET("/contact/delete", func(c *gin.Context) { appController.Contact.Delete(&adapter.ContextAdapter{Context: c}) })
	router.GET("/contact/info", func(c *gin.Context) { appController.Contact.Info(&adapter.ContextAdapter{Context: c}) })

	router.GET("/transfer/check", func(c *gin.Context) { appController.Transfer.Check(&adapter.ContextAdapter{Context: c}) })
	router.GET("/transfer/request", func(c *gin.Context) { appController.Transfer.Request(&adapter.ContextAdapter{Context: c}) })
	router.GET("/transfer/cancel", func(c *gin.Context) { appController.Transfer.Cancel(&adapter.ContextAdapter{Context: c}) })
	router.GET("/transfer/approve", func(c *gin.Context) { appController.Transfer.Approve(&adapter.ContextAdapter{Context: c}) })
	router.GET("/transfer/reject", func(c *gin.Context) { appController.Transfer.Approve(&adapter.ContextAdapter{Context: c}) })

	router.GET("/poll", func(c *gin.Context) { appController.Poll.Poll(&adapter.ContextAdapter{Context: c}) })

	return router
}
