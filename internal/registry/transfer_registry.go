package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/interface/controller"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewTransferController() controller.TransferController {
	registrarInteractor := interactor.NewTransferInteractor(
		repository.NewRegistrarRepository(r.eppClient),
		presenter.NewTransferPresenter(),
	)

	return controller.NewTransferController(registrarInteractor)
}
