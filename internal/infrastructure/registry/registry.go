package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter"
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/delivery/http/controllers"
	usecase "gitlab.com/merekmu/go-epp-rest/internal/usecase/infrastructure"
	"gorm.io/gorm"
)

type registry struct {
	eppClient adapter.EppClient
	mysqlConn *gorm.DB
	xmlMapper usecase.XMLMapper
}

type Registry interface {
	NewAppController() controllers.AppController
}

func NewRegistry(eppClient adapter.EppClient, mysqlConn *gorm.DB, xmlMapper usecase.XMLMapper) Registry {
	return &registry{eppClient: eppClient, mysqlConn: mysqlConn, xmlMapper: xmlMapper}
}

func (r *registry) NewAppController() controllers.AppController {
	return controllers.AppController{
		Domain:   r.NewDomainController(),
		Contact:  r.NewContactController(),
		Host:     r.NewHostController(),
		Transfer: r.NewTransferController(),
		Poll:     r.NewPollController(),
	}
}
