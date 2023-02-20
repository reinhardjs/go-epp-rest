package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type domainPresenter struct{}

func NewUserPresenter() presenter.DomainPresenter {
	return &domainPresenter{}
}

func (up *domainPresenter) ResponseQueryDomain(response string) (string, error) {
	return response, nil
}
