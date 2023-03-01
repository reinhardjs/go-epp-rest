package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/delivery"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure"
	"gorm.io/gorm"
)

type registry struct {
	eppClient infrastructure.EppClient
	mysqlConn *gorm.DB
}

type Registry interface {
	NewAppController() delivery.AppController
}

func NewRegistry(eppClient infrastructure.EppClient, mysqlConn *gorm.DB) Registry {
	return &registry{eppClient: eppClient, mysqlConn: mysqlConn}
}

func (r *registry) NewAppController() delivery.AppController {
	return delivery.AppController{
		Domain:   r.NewDomainController(),
		Contact:  r.NewContactController(),
		Host:     r.NewHostController(),
		Transfer: r.NewTransferController(),
	}
}
