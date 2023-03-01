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
	Cancel(data interface{}, ext string, langTag string) (res string, err error)
	Approve(data interface{}, ext string, langTag string) (res string, err error)
	Reject(data interface{}, ext string, langTag string) (res string, err error)
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

	responseObj, err := interactor.Presenter.Check(responseByte)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Check: interactor.Presenter.MapCheckResponse")
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

	responseObj, err := interactor.Presenter.Request(responseByte)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.Presenter.MapRequestResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}

func (interactor *transferInteractor) Cancel(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Cancel(responseByte)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.Presenter.MapCancelResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}

func (interactor *transferInteractor) Approve(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Approve(responseByte)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.Presenter.MapApproveResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}

func (interactor *transferInteractor) Reject(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Reject(responseByte)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.Presenter.MapRejectResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}
