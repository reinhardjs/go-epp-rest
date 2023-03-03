package interactor

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/common/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type domainInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.DomainPresenter
	XMLMapper           adapter.XMLMapper
}

type DomainInteractor interface {
	Check(data interface{}, ext string, langTag string) (res string, err error)
	Create(data interface{}, ext string, langTag string) (res string, err error)
	Delete(data interface{}, ext string, langTag string) (res string, err error)
	Info(data interface{}, ext string, langTag string) (res string, err error)
	SecDNSUpdate(data interface{}, ext string, langTag string) (res string, err error)
}

func NewDomainInteractor(domainRepository repository.RegistrarRepository, presenter presenter.DomainPresenter, xmlMapper adapter.XMLMapper) DomainInteractor {
	return &domainInteractor{
		RegistrarRepository: domainRepository,
		Presenter:           presenter,
		XMLMapper:           xmlMapper,
	}
}

func (interactor *domainInteractor) Check(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Check: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CheckDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	res = interactor.Presenter.Check(*responseDTO)
	return
}

func (interactor *domainInteractor) Create(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CreateDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	res = interactor.Presenter.Create(*responseDTO)
	return
}

func (interactor *domainInteractor) Delete(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.DeleteDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	res = interactor.Presenter.Delete(*responseDTO)
	return
}

func (interactor *domainInteractor) Info(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.InfoDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	res = interactor.Presenter.Info(*responseDTO)
	return
}

func (interactor *domainInteractor) SecDNSUpdate(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.SecDNSUpdateResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	res = interactor.Presenter.SecDNSUpdate(*responseDTO)
	return
}
