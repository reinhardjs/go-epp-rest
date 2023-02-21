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

type hostInteractor[T constraints.RegistrarResponseConstraint] struct {
	RegistrarRepository repository.RegistrarRepository
	RegistrarPresenter  presenter.RegistrarPresenter[T]
}

func NewHostInteractor[T constraints.RegistrarResponseConstraint](repository repository.RegistrarRepository, presenter presenter.RegistrarPresenter[T]) RegistrarInteractor {
	return &hostInteractor[T]{
		RegistrarRepository: repository,
		RegistrarPresenter:  presenter,
	}
}

func (interactor *hostInteractor[T]) Check(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.Check(data)
	if err != nil {
		err = errors.Wrap(err, "Host Interactor: Check interactor.RegistrarRepository.Check")
		return
	}

	log.Println("XML Response: \n", string(responseByte))

	genericResponseObj, err := interactor.RegistrarPresenter.Check(responseByte)

	if err != nil {
		err = errors.Wrap(err, "Host Interactor: Check interactor.RegistrarPresenter.Check")
		return
	}

	// converting from generic object into model object
	responseObj := any(genericResponseObj).(model.CheckHostResponse)

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
