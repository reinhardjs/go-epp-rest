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

type contactInteractor[T constraints.RegistrarResponseConstraint] struct {
	RegistrarRepository repository.RegistrarRepository
	RegistrarPresenter  presenter.RegistrarPresenter[T]
}

type ContactInteractor interface {
	Check(data interface{}, ext string, langTag string) (string, error)
}

func NewContactInteractor[T constraints.RegistrarResponseConstraint](repository repository.RegistrarRepository, presenter presenter.RegistrarPresenter[T]) ContactInteractor {
	return &contactInteractor[T]{
		RegistrarRepository: repository,
		RegistrarPresenter:  presenter,
	}
}

func (interactor *contactInteractor[T]) Check(data interface{}, ext string, langTag string) (res string, returnedErr error) {
	response, err := interactor.RegistrarRepository.Check(data)
	if err != nil {
		returnedErr = errors.Wrap(err, "Contact Interactor: Check interactor.RegistrarRepository.Check")
		return
	}

	log.Println("XML Response: \n", string(response))

	genericResponseObj, err := interactor.RegistrarPresenter.ResponseCheck(response)

	if err != nil {
		returnedErr = errors.Wrap(err, "Contact Interactor: Check interactor.RegistrarPresenter.ResponseCheck")
		return
	}

	// converting from generic object into model object
	responseObj := any(genericResponseObj).(model.CheckContactResponse)

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
