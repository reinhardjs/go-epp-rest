package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewHostController() controllers.HostController {
	registrarInteractor := interactor.NewHostInteractor(
		repository.NewRegistrarRepository(r.eppClient),
		presenter.NewHostPresenter(),
	)

	return controllers.NewHostController(registrarInteractor)
}
