package registry

import (
	"gitlab.com/merekmu/go-epp-rest/internal/adapter/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

func (r *registry) NewPollController() controllers.PollController {
	pollInteractor := interactor.NewPollInteractor(
		repository.NewEppPollRepository(r.mysqlConn),
		repository.NewRegistrarRepository(r.eppClient, r.xmlMapper),
		presenter.NewPollPresenter(),
		r.xmlMapper,
		r.dtoToEntityMapper,
	)

	return controllers.NewPollController(pollInteractor)
}
