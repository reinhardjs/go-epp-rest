package interactor

import (
	"fmt"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type pollInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.PollPresenter
}

type PollInteractor interface {
	Poll() (res string, err error)
}

func NewPollInteractor(repository repository.RegistrarRepository, presenter presenter.PollPresenter) PollInteractor {
	return &pollInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
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

	responseObj, err := interactor.Presenter.Request(responseByte)

	if err != nil {
		err = errors.Wrap(err, "PollInteractor Request: interactor.Presenter.MapRequestResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}
