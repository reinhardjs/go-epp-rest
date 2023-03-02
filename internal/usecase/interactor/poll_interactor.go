package interactor

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/adapter/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/repository"
	infrastructure "gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type pollInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.PollPresenter
	XMLMapper           infrastructure.XMLMapper
}

type PollInteractor interface {
	Poll() (res string, err error)
}

func NewPollInteractor(repository repository.RegistrarRepository, presenter presenter.PollPresenter, xmlMapper infrastructure.XMLMapper) PollInteractor {
	return &pollInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
		XMLMapper:           xmlMapper,
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

	responseObj := &response.PollRequestResponse{}
	err = interactor.XMLMapper.MapXMLToModel(responseByte, responseObj)
	res = interactor.Presenter.Request(responseObj)
	return
}
