package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/interface/controller"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewDomainController() controller.DomainController {
	registrarInteractor := interactor.NewDomainInteractor(
		repository.NewRegistrarRepository(r.eppClient),
		presenter.NewRegistrarPresenter(),
	)

	return controller.NewDomainController(registrarInteractor)
}
