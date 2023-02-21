package repository

import (
	"github.com/bombsimon/epp-go"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type registrarRepository struct {
	eppClient infrastructure.EppClient
}

func NewRegistrarRepository(eppClient infrastructure.EppClient) repository.RegistrarRepository {
	return &registrarRepository{eppClient}
}

func (r *registrarRepository) sendXMLTCPRequest(data interface{}) ([]byte, error) {
	encoded, err := epp.Encode(data, epp.ClientXMLAttributes())
	if err != nil {
		return nil, errors.Wrap(err, "registrarRepository sendXMLTCPRequest: epp.Encode")
	}

	byteResponse, err := r.eppClient.Send(encoded)
	if err != nil {
		return nil, errors.Wrap(err, "registrarRepository sendXMLTCPRequest: r.eppClient.Send")
	}

	return byteResponse, nil
}

func (r *registrarRepository) Check(data interface{}) ([]byte, error) {
	return r.sendXMLTCPRequest(data)
}
