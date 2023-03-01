package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/delivery"
	"gitlab.com/merekmu/go-epp-rest/internal/presentation"
	"gitlab.com/merekmu/go-epp-rest/internal/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewContactController() delivery.ContactController {
	registrarInteractor := interactor.NewContactInteractor(
		repository.NewRegistrarRepository(r.eppClient),
		presentation.NewContactPresenter(),
	)

	return delivery.NewContactController(registrarInteractor)
}
