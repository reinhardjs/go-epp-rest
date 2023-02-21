package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type registrarPresenter struct{}

func NewRegistrarPresenter() presenter.RegistrarPresenter {
	return &registrarPresenter{}
}

func (up *registrarPresenter) ResponseCheck(response string) (string, error) {
	return response, nil
}
