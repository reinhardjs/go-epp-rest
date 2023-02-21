package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/controller"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewHostController() controller.HostController {
	registrarInteractor := interactor.NewHostInteractor(
		repository.NewRegistrarRepository(r.eppClient),
		presenter.NewRegistrarPresenter[model.CheckHostResponse](),
	)

	return controller.NewHostController(registrarInteractor)
}
