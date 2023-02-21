package interactor

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/constraints"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type domainInteractor[T constraints.RegistrarResponseConstraint] struct {
	RegistrarRepository repository.RegistrarRepository
	RegistrarPresenter  presenter.RegistrarPresenter[T]
}

func NewDomainInteractor[T constraints.RegistrarResponseConstraint](domainRepository repository.RegistrarRepository, domainPresenter presenter.RegistrarPresenter[T]) RegistrarInteractor[T] {
	return &domainInteractor[T]{
		RegistrarRepository: domainRepository,
		RegistrarPresenter:  domainPresenter,
	}
}

func (interactor *domainInteractor[T]) Send(data interface{}) (res T, err error) {
	responseByte, err := interactor.RegistrarRepository.Check(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Send: interactor.RegistrarRepository.Check")
		return
	}

	log.Println("XML Response: \n", string(responseByte))

	genericResponseObj, err := interactor.RegistrarPresenter.Check(responseByte)

	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Send: interactor.RegistrarPresenter.Check")
		return
	}

	res = genericResponseObj
	return
}

func (interactor *domainInteractor[T]) Check(data interface{}, ext string, langTag string) (res string, returnedErr error) {
	genericResponseObj, err := interactor.Send(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Check: interactor.Send")
		return
	}

	// converting from generic object into model object
	modelResponseObj := any(genericResponseObj).(model.CheckDomainResponse)

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
