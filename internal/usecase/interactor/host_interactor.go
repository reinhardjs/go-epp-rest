package interactor

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type hostInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.HostPresenter
}

type HostInteractor interface {
	Check(data interface{}, ext string, langTag string) (res string, err error)
}

func NewHostInteractor(repository repository.RegistrarRepository, presenter presenter.HostPresenter) HostInteractor {
	return &hostInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
	}
}

func (interactor *hostInteractor) Check(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Send: interactor.RegistrarRepository.SendCommand")
		return
	}

	log.Println("XML Response: \n", string(responseByte))

	responseObj, err := interactor.Presenter.MapCheckResponse(responseByte)

	if err != nil {
		err = errors.Wrap(err, "HostInteractor Send: interactor.HostPresenter.MapResponse")
		return
	}

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
