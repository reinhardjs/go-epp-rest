package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/adapter/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewTransferController() controllers.TransferController {
	registrarInteractor := interactor.NewTransferInteractor(
		repository.NewRegistrarRepository(r.eppClient, r.xmlMapper),
		presenter.NewTransferPresenter(),
	)

	return controllers.NewTransferController(registrarInteractor)
}
