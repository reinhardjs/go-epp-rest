package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/delivery"
	"gitlab.com/merekmu/go-epp-rest/internal/presentation"
	"gitlab.com/merekmu/go-epp-rest/internal/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewDomainController() delivery.DomainController {
	registrarInteractor := interactor.NewDomainInteractor(
		repository.NewRegistrarRepository(r.eppClient),
		presentation.NewDomainPresenter(),
	)

	return delivery.NewDomainController(registrarInteractor)
}
