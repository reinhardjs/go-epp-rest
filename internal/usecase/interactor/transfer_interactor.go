package interactor

import (
	"fmt"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type transferInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.TransferPresenter
}

type TransferInteractor interface {
	Check(data interface{}, ext string, langTag string) (res string, err error)
	Request(data interface{}, ext string, langTag string) (res string, err error)
}

func NewTransferInteractor(repository repository.RegistrarRepository, presenter presenter.TransferPresenter) TransferInteractor {
	return &transferInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
	}
}

func (interactor *transferInteractor) Check(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Check: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.MapCheckResponse(responseByte)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Check: interactor.TransferPresenter.MapCheckResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}

func (interactor *transferInteractor) Request(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.MapRequestResponse(responseByte)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.TransferPresenter.MapRequestResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}
