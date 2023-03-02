package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewPollController() controllers.PollController {
	pollInteractor := interactor.NewPollInteractor(
		repository.NewRegistrarRepository(r.eppClient),
		presenter.NewPollPresenter(),
		r.xmlMapper,
	)

	return controllers.NewPollController(pollInteractor)
}
