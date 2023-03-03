package interactor

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/usecase"
)

type contactInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.ContactPresenter
}

func NewContactInteractor(repository repository.RegistrarRepository, presenter presenter.ContactPresenter) usecase.ContactInteractor {
	return &contactInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
	}
}

func (interactor *contactInteractor) Check(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Check: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Check(responseByte)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Check: interactor.Presenter.MapCheckResponse")
		return
	}

	for _, element := range responseObj.ResultData.CheckDatas {
		notStr := ""
		if element.Id.AvailKey == 0 {
			notStr = "not "
		}
		res += fmt.Sprintf("Contact %s, contact %savailable\n", element.Id.Value, notStr)
	}
	res = strings.TrimSuffix(res, "\n")

	return
}

func (interactor *contactInteractor) Create(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Create(responseByte)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Create: interactor.Presenter.MapCreateResponse")
		return
	}

	res += fmt.Sprintf("ID %s\n", responseObj.ResultData.CreateData.Id)
	res += fmt.Sprintf("Create Date %s\n", responseObj.ResultData.CreateData.CreateDate)
	res = strings.TrimSuffix(res, "\n")

	return
}

func (interactor *contactInteractor) Update(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Update: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Update(responseByte)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Update: interactor.Presenter.MapCreateResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}

func (interactor *contactInteractor) Delete(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Delete: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Delete(responseByte)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Delete: interactor.Presenter.MapDeleteResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}

func (interactor *contactInteractor) Info(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Info: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Info(responseByte)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Info: interactor.Presenter.MapInfoResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}
