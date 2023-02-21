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

type contactInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.ContactPresenter
}

type ContactInteractor interface {
	Send(data interface{}) (interface{}, error)
	Check(data interface{}, ext string, langTag string) (res string, err error)
}

func NewContactInteractor(repository repository.RegistrarRepository, presenter presenter.ContactPresenter) ContactInteractor {
	return &contactInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
	}
}

func (interactor *contactInteractor) Send(data interface{}) (res interface{}, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Send: interactor.RegistrarRepository.SendCommand")
		return
	}

	log.Println("XML Response: \n", string(responseByte))

	genericResponseObj, err := interactor.Presenter.MapResponse(responseByte)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Send: interactor.ContactPresenter.MapResponse")
		return
	}

	res = genericResponseObj
	return
}

func (interactor *contactInteractor) Check(data interface{}, ext string, langTag string) (res string, returnedErr error) {
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
