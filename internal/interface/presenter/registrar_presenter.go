package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type domainPresenter struct{}

func NewUserPresenter() presenter.RegistrarPresenter {
	return &domainPresenter{}
}

func (up *domainPresenter) ResponseCheck(response string) (string, error) {
	return response, nil
}
