package interactor

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type domainInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.DomainPresenter
}

type DomainInteractor interface {
	Send(data interface{}) (interface{}, error)
	Check(data interface{}, ext string, langTag string) (res string, err error)
}

func NewDomainInteractor(domainRepository repository.RegistrarRepository, presenter presenter.DomainPresenter) DomainInteractor {
	return &domainInteractor{
		RegistrarRepository: domainRepository,
		Presenter:           presenter,
	}
}

func (interactor *domainInteractor) Send(data interface{}) (res interface{}, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Send: interactor.RegistrarRepository.SendCommand")
		return
	}

	log.Println("XML Response: \n", string(responseByte))

	genericResponseObj, err := interactor.Presenter.MapResponse(responseByte)

	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Send: interactor.DomainPresenter.MapResponse")
		return
	}

	res = genericResponseObj
	return
}

func (interactor *domainInteractor) Check(data interface{}, ext string, langTag string) (res string, returnedErr error) {
	genericResponseObj, err := interactor.Send(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Check: interactor.Send")
		return
	}

	log.Println(genericResponseObj)

	// converting from generic object into model object
	modelResponseObj := genericResponseObj.(model.CheckDomainResponse)

	for _, element := range modelResponseObj.ResultData.CheckDatas {
		notStr := ""
		if element.Name.AvailKey == 0 {
			notStr = "not "
		}
		res += fmt.Sprintf("Domain %s, domain %savailable\n", element.Name.Value, notStr)
	}
	res = strings.TrimSuffix(res, "\n")

	return
}
