package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/delivery"
	"gitlab.com/merekmu/go-epp-rest/internal/presentation"
	"gitlab.com/merekmu/go-epp-rest/internal/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewHostController() delivery.HostController {
	registrarInteractor := interactor.NewHostInteractor(
		repository.NewRegistrarRepository(r.eppClient),
		presentation.NewHostPresenter(),
	)

	return delivery.NewHostController(registrarInteractor)
}
