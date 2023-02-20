package interactor

import (
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type domainInteractor struct {
	DomainRepository repository.DomainRepository
	DomainPresenter  presenter.DomainPresenter
}

type DomainInteractor interface {
	DoQueryDomain(domainList []string, ext string, langTag string) (string, error)
}

func NewDomainInteractor(domainRepository repository.DomainRepository, domainPresenter presenter.DomainPresenter) DomainInteractor {
	return &domainInteractor{domainRepository, domainPresenter}
}

func (interactor *domainInteractor) DoQueryDomain(domainList []string, ext string, langTag string) (string, error) {
	response, err := interactor.DomainRepository.DoQueryDomain(domainList, ext, langTag)

	if err != nil {
		return "", err
	}

	return interactor.DomainPresenter.ResponseQueryDomain(response)
}
