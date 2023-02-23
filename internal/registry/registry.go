package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/controller"
)

type registry struct {
	eppClient infrastructure.EppClient
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(eppClient infrastructure.EppClient) Registry {
	return &registry{eppClient: eppClient}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Domain:   r.NewDomainController(),
		Contact:  r.NewContactController(),
		Host:     r.NewHostController(),
		Transfer: r.NewTransferController(),
	}
}
