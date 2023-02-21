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

func NewContactInteractor[T constraints.RegistrarResponseConstraint](repository repository.RegistrarRepository, presenter presenter.RegistrarPresenter[T]) RegistrarInteractor[T] {
	return &contactInteractor[T]{
		RegistrarRepository: repository,
		RegistrarPresenter:  presenter,
	}
}

func (interactor *contactInteractor[T]) Send(data interface{}) (res T, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Send: interactor.RegistrarRepository.SendCommand")
		return
	}

	log.Println("XML Response: \n", string(responseByte))

	genericResponseObj, err := interactor.RegistrarPresenter.MapResponse(responseByte)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Send: interactor.RegistrarPresenter.MapResponse")
		return
	}

	res = genericResponseObj
	return
}

func (interactor *contactInteractor[T]) Check(data interface{}, ext string, langTag string) (res string, returnedErr error) {
	genericResponseObj, err := interactor.Send(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Check: interactor.Send")
		return
	}

	// converting from generic object into model object
	modelResponseObj := any(genericResponseObj).(model.CheckContactResponse)

	for _, element := range modelResponseObj.ResultData.CheckDatas {
		notStr := ""
		if element.Id.AvailKey == 0 {
			notStr = "not "
		}
		res += fmt.Sprintf("Contact %s, contact %savailable\n", element.Id.Value, notStr)
	}
	res = strings.TrimSuffix(res, "\n")

	return
}
