package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/adapter/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewDomainController() controllers.DomainController {
	registrarInteractor := interactor.NewDomainInteractor(
		repository.NewRegistrarRepository(r.eppClient, r.xmlMapper),
		presenter.NewDomainPresenter(),
		r.xmlMapper,
	)

	return controllers.NewDomainController(registrarInteractor)
}
