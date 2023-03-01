package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/delivery/http/controllers"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewTransferController() controllers.TransferController {
	registrarInteractor := interactor.NewTransferInteractor(
		repository.NewRegistrarRepository(r.eppClient),
		presenter.NewTransferPresenter(),
	)

	return controllers.NewTransferController(registrarInteractor)
}
