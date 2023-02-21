package interactor

import (
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type registrarInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	RegistrarPresenter  presenter.RegistrarPresenter
}

type RegistrarInteractor interface {
	Check(data interface{}, ext string, langTag string) (string, error)
}

func NewDomainInteractor(registrarRepository repository.RegistrarRepository, registrarPresenter presenter.RegistrarPresenter) RegistrarInteractor {
	return &registrarInteractor{
		RegistrarRepository: registrarRepository,
		RegistrarPresenter:  registrarPresenter,
	}
}

func (interactor *registrarInteractor) Check(data interface{}, ext string, langTag string) (string, error) {
	response, err := interactor.RegistrarRepository.Check(data)

	if err != nil {
		return "", err
	}

	return interactor.RegistrarPresenter.ResponseCheck(response)
}
