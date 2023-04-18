package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter/mapper"
	"gorm.io/gorm"
)

type registry struct {
	eppClient         adapter.EppClient
	mysqlConn         *gorm.DB
	xmlMapper         mapper.XMLMapper
	dtoToEntityMapper mapper.DtoToEntity
}

type Registry interface {
	NewAppController() controllers.AppController
}

func NewRegistry(eppClient adapter.EppClient, mysqlConn *gorm.DB, xmlMapper mapper.XMLMapper, dtoToEntityMapper mapper.DtoToEntity) Registry {
	return &registry{eppClient: eppClient, mysqlConn: mysqlConn, xmlMapper: xmlMapper, dtoToEntityMapper: dtoToEntityMapper}
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
