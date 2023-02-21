package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/controller"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewContactController() controller.ContactController {
	registrarInteractor := interactor.NewContactInteractor(
		repository.NewRegistrarRepository(r.eppClient),
		presenter.NewRegistrarPresenter[model.ContactCheckResponse](),
	)

	return controller.NewContactController(registrarInteractor)
}
