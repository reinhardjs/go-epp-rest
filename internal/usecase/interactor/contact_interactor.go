package interactor

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type contactInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.ContactPresenter
}

type ContactInteractor interface {
	Check(data interface{}, ext string, langTag string) (res string, err error)
}

func NewContactInteractor(repository repository.RegistrarRepository, presenter presenter.ContactPresenter) ContactInteractor {
	return &contactInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
	}
}

func (interactor *contactInteractor) Check(data interface{}, ext string, langTag string) (res string, returnedErr error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Send: interactor.RegistrarRepository.SendCommand")
		return
	}

	log.Println("XML Response: \n", string(responseByte))

	responseObj, err := interactor.Presenter.MapCheckResponse(responseByte)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Send: interactor.ContactPresenter.MapResponse")
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
