package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/delivery"
	"gitlab.com/merekmu/go-epp-rest/internal/presentation"
	"gitlab.com/merekmu/go-epp-rest/internal/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewTransferController() delivery.TransferController {
	registrarInteractor := interactor.NewTransferInteractor(
		repository.NewRegistrarRepository(r.eppClient),
		presentation.NewTransferPresenter(),
	)

	return delivery.NewTransferController(registrarInteractor)
}
