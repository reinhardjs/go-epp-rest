package interactor

import (
	"fmt"
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

type DomainInteractor interface {
	Check(data interface{}, ext string, langTag string) (string, error)
}

func NewDomainInteractor[T constraints.RegistrarResponseConstraint](domainRepository repository.RegistrarRepository, domainPresenter presenter.RegistrarPresenter[T]) DomainInteractor {
	return &domainInteractor[T]{
		RegistrarRepository: domainRepository,
		RegistrarPresenter:  domainPresenter,
	}
}

func (interactor *domainInteractor[T]) Check(data interface{}, ext string, langTag string) (res string, returnedErr error) {
	response, err := interactor.RegistrarRepository.Check(data)
	if err != nil {
		returnedErr = errors.Wrap(err, "Domain Interactor: Check interactor.RegistrarRepository.Check")
		return
	}

	genericResponseObj, err := interactor.RegistrarPresenter.ResponseCheck(response)

	if err != nil {
		returnedErr = errors.Wrap(err, "Domain Interactor: Check interactor.RegistrarPresenter.ResponseCheck")
		return
	}

	// converting from generic object into model object
	responseObj := any(genericResponseObj).(model.CheckDomainResponse)

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
