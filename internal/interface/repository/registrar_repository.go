package repository

import (
	"github.com/bombsimon/epp-go"
	"github.com/bombsimon/epp-go/types"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type domainRepository struct {
	eppClient infrastructure.EppClient
}

func NewDomainRepository(eppClient infrastructure.EppClient) repository.RegistrarRepository {
	return &domainRepository{eppClient}
}

func (r *domainRepository) Check(list []string, ext string, langTag string) (string, error) {
	domainCheck := types.DomainCheckType{
		Check: types.DomainCheck{
			Names: list,
		},
	}

	encoded, err := epp.Encode(domainCheck, epp.ClientXMLAttributes())
	if err != nil {
		return "", err
	}

	byteResponse, err := r.eppClient.Send(encoded)
	if err != nil {
		return "", err
	}

	return string(byteResponse), nil
}
