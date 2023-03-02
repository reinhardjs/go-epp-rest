package interactor

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type domainInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.DomainPresenter
}

type DomainInteractor interface {
	Check(data interface{}, ext string, langTag string) (res string, err error)
	Create(data interface{}, ext string, langTag string) (res string, err error)
	Delete(data interface{}, ext string, langTag string) (res string, err error)
	Info(data interface{}, ext string, langTag string) (res string, err error)
	SecDNSUpdate(data interface{}, ext string, langTag string) (res string, err error)
}

func NewDomainInteractor(domainRepository repository.RegistrarRepository, presenter presenter.DomainPresenter) DomainInteractor {
	return &domainInteractor{
		RegistrarRepository: domainRepository,
		Presenter:           presenter,
	}
}

func (interactor *domainInteractor) Check(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Check: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Check(responseByte)

	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Check: interactor.Presenter.MapResponse")
		return
	}

	for _, element := range responseObj.ResultData.CheckDatas {
		notStr := ""
		if element.Name.AvailKey == 0 {
			notStr = "not "
		}
		res += fmt.Sprintf("Domain %s, domain %savailable\n", element.Name.Value, notStr)
	}
	res = strings.TrimSuffix(res, "\n")

	return
}

func (interactor *domainInteractor) Create(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Create(responseByte)

	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.Presenter.MapCreateResponse")
		return
	}

	res += fmt.Sprintf("Name %s\n", responseObj.ResultData.CreatedData.Name)
	res += fmt.Sprintf("Create Date %s\n", responseObj.ResultData.CreatedData.CreatedDate)
	res += fmt.Sprintf("Expire Date %s\n", responseObj.ResultData.CreatedData.ExpiredDate)
	res = strings.TrimSuffix(res, "\n")

	return
}

func (interactor *domainInteractor) Delete(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Delete: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Delete(responseByte)

	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Delete: interactor.Presenter.MapDeleteResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}

func (interactor *domainInteractor) Info(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Info: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseObj, err := interactor.Presenter.Info(responseByte)

	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Info: interactor.Presenter.MapInfoResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}

func (interactor *domainInteractor) SecDNSUpdate(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor SecDNSUpdate: interactor.RegistrarRepository.SecDNSUpdate")
		return
	}

	responseObj, err := interactor.Presenter.SecDNSUpdate(responseByte)

	if err != nil {
		err = errors.Wrap(err, "DomainInteractor SecDNSUpdate: interactor.Presenter.MapSecDNSUpdateResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result.Code, responseObj.Result.Message)

	return
}
