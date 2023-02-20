package repository

import (
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type domainRepository struct {
}

func NewDomainRepository() repository.DomainRepository {
	return &domainRepository{}
}

func (r *domainRepository) DoQueryDomain(domainList []string, ext string, langTag string) (string, error) {
	return "", nil
}
