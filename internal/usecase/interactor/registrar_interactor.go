package interactor

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/constraints"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type registrarInteractor[T constraints.RegistrarResponseConstraint] struct {
	RegistrarRepository repository.RegistrarRepository
	RegistrarPresenter  presenter.RegistrarPresenter[T]
}

type RegistrarInteractor interface {
	Check(data interface{}, ext string, langTag string) (string, error)
}

func NewRegistrarInteractor[T constraints.RegistrarResponseConstraint](registrarRepository repository.RegistrarRepository, registrarPresenter presenter.RegistrarPresenter[T]) RegistrarInteractor {
	return &registrarInteractor[T]{
		RegistrarRepository: registrarRepository,
		RegistrarPresenter:  registrarPresenter,
	}
}

func (interactor *registrarInteractor[T]) Check(data interface{}, ext string, langTag string) (string, error) {
	response, err := interactor.RegistrarRepository.Check(data)

	if err != nil {
		return "", errors.Wrap(err, "Domain Interactor: Check interactor.RegistrarRepository.Check")
	}

	return interactor.RegistrarPresenter.ResponseCheck(response)
}
