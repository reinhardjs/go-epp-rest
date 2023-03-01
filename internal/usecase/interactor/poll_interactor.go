package interactor

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type pollInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.PollPresenter
	xmlMapper           infrastructure.XMLMapper
}

type PollInteractor interface {
	Poll() (res string, err error)
}

func NewPollInteractor(repository repository.RegistrarRepository, presenter presenter.PollPresenter, xmlMapper infrastructure.XMLMapper) PollInteractor {
	return &pollInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
		xmlMapper:           xmlMapper,
	}
}

func (interactor *pollInteractor) Poll() (res string, err error) {
	data := types.Poll{
		Poll: types.PollCommand{
			Operation: types.PollOperationRequest,
		},
	}

	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "PollInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.xmlMapper.MapXMLToModel(string(responseByte))
	res = interactor.Presenter.Request(responseObj)
	return
}
