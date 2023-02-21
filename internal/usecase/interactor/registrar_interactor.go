package interactor

import (
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type domainInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	RegistrarPresenter  presenter.RegistrarPresenter
}

type DomainInteractor interface {
	Check(list []string, ext string, langTag string) (string, error)
}

func NewDomainInteractor(registrarRepository repository.RegistrarRepository, registrarPresenter presenter.RegistrarPresenter) DomainInteractor {
	return &domainInteractor{
		RegistrarRepository: registrarRepository,
		RegistrarPresenter:  registrarPresenter,
	}
}

func (interactor *domainInteractor) Check(list []string, ext string, langTag string) (string, error) {
	response, err := interactor.RegistrarRepository.Check(list, ext, langTag)

	if err != nil {
		return "", err
	}

	return interactor.RegistrarPresenter.ResponseCheck(response)
}
