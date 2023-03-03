package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewContactController() controllers.ContactController {
	registrarInteractor := interactor.NewContactInteractor(
		repository.NewRegistrarRepository(r.eppClient, r.xmlMapper),
		presenter.NewContactPresenter(),
	)

	return controllers.NewContactController(registrarInteractor)
}
