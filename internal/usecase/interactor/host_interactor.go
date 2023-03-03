package interactor

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type hostInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.HostPresenter
}

func NewHostInteractor(repository repository.RegistrarRepository, presenter presenter.HostPresenter) usecase.HostInteractor {
	return &hostInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
	}
}

func (interactor *hostInteractor) Check(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Send: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Check(responseByte)

	if err != nil {
		err = errors.Wrap(err, "HostInteractor Send: interactor.HostPresenter.MapResponse")
		return
	}

	for _, element := range responseObj.ResultData.CheckDatas {
		notStr := ""
		if element.HostName.AvailKey == 0 {
			notStr = "not "
		}
		res += fmt.Sprintf("Host %s, host %savailable\n", element.HostName.Value, notStr)
	}
	res = strings.TrimSuffix(res, "\n")

	return
}

func (interactor *hostInteractor) Create(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Create(responseByte)

	if err != nil {
		err = errors.Wrap(err, "HostInteractor Create: interactor.Presenter.MapCreateResponse")
		return
	}

	res += fmt.Sprintf("Name %s\n", responseObj.ResultData.CreateData.Name)
	res += fmt.Sprintf("Create Date %s\n", responseObj.ResultData.CreateData.CreateDate)
	res = strings.TrimSuffix(res, "\n")

	return
}

func (interactor *hostInteractor) Update(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Update: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Update(responseByte)

	if err != nil {
		err = errors.Wrap(err, "HostInteractor Update: interactor.Presenter.MapUpdateResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}

func (interactor *hostInteractor) Delete(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Delete: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Delete(responseByte)

	if err != nil {
		err = errors.Wrap(err, "HostInteractor Delete: interactor.Presenter.MapDeleteResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}

func (interactor *hostInteractor) Info(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Info: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Info(responseByte)

	if err != nil {
		err = errors.Wrap(err, "HostInteractor Info: interactor.Presenter.MapInfoResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}
