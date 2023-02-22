package interactor

import (
	"fmt"
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
	Create(data interface{}, ext string, langTag string) (res string, err error)
	Update(data interface{}, ext string, langTag string) (res string, err error)
}

func NewContactInteractor(repository repository.RegistrarRepository, presenter presenter.ContactPresenter) ContactInteractor {
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

	responseObj, err := interactor.Presenter.MapCheckResponse(responseByte)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Check: interactor.ContactPresenter.MapCheckResponse")
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

	responseObj, err := interactor.Presenter.MapCreateResponse(responseByte)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Create: interactor.ContactPresenter.MapCreateResponse")
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

	responseObj, err := interactor.Presenter.MapUpdateResponse(responseByte)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Update: interactor.ContactPresenter.MapCreateResponse")
		return
	}

	res = fmt.Sprintf("%v %v", responseObj.Result[0].Code, responseObj.Result[0].Message)

	return
}
