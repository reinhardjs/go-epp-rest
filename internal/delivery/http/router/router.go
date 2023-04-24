package router

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/middlewares"

	"github.com/gin-contrib/pprof"
)

func NewRouter(appController controllers.AppController) *gin.Engine {
	handler := &handler{appController}
	router := gin.Default()

	// Profiling monitor
	pprof.Register(router)

	// Prometheus monitoring
	prometheus := ginprometheus.NewPrometheus("gin")
	prometheus.Use(router)

	// Use error handler middleware
	router.Use(middlewares.ClientErrorHandler)                     // Error related to client error
	router.Use(gin.CustomRecovery(middlewares.ServerErrorHandler)) // Error related to server error resulted from like panic/exception, etc..

	v1 := router.Group("api/v1")
	{
		v1.GET("/", handler.apiV1)
	}

	v2 := router.Group("api/v2")
	{
		v2.GET("/domain/check", handler.domainCheck)
		v2.GET("/domain/create", handler.domainCreate)
		v2.GET("/domain/delete", handler.domainDelete)
		v2.GET("/domain/info", handler.domainInfo)
		v2.GET("/domain/secdnsupdate", handler.domainSecDNSUpdate)
		v2.GET("/domain/contact/update", handler.domainContactUpdate)
		v2.GET("/domain/status/update", handler.domainStatusUpdate)
		v2.GET("/domain/authinfo/update", handler.domainAuthInfoUpdate)
		v2.GET("/domain/nameserver/update", handler.domainNameserverUpdate)
		v2.GET("/domain/renew", handler.domainRenew)

		v2.GET("/host/check", handler.hostCheck)
		v2.GET("/host/create", handler.hostCreate)
		v2.GET("/host/update", handler.hostUpdate)
		v2.GET("/host/delete", handler.hostDelete)
		v2.GET("/host/info", handler.hostInfo)
		v2.GET("/host/change", handler.hostChange)
		v2.GET("/host/checkcreate", handler.hostCheckAndCreate)

		v2.GET("/contact/check", handler.contactCheck)
		v2.GET("/contact/create", handler.contactCreate)
		v2.GET("/contact/update", handler.contactUpdate)
		v2.GET("/contact/delete", handler.contactDelete)
		v2.GET("/contact/info", handler.contactInfo)

		v2.GET("/transfer/check", handler.transferCheck)
		v2.GET("/transfer/request", handler.transferRequest)
		v2.GET("/transfer/cancel", handler.transferCancel)
		v2.GET("/transfer/approve", handler.transferApprove)
		v2.GET("/transfer/reject", handler.transferReject)

		v2.GET("/poll", handler.poll)
	}

	return router
}
