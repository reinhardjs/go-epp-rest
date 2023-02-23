package repository

import (
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp"
)

type registrarRepository struct {
	eppClient infrastructure.EppClient
}

func NewRegistrarRepository(eppClient infrastructure.EppClient) repository.RegistrarRepository {
	return &registrarRepository{eppClient}
}

func (r *registrarRepository) prepareCommand(data interface{}) ([]byte, error) {
	encoded, err := registry_epp.Encode(data, registry_epp.ClientXMLAttributes())
	if err != nil {
		return nil, errors.Wrap(err, "registrarRepository prepareCommand: registry_epp.Encode")
	}

	return encoded, nil
}

func (r *registrarRepository) SendCommand(data interface{}) ([]byte, error) {
	encoded, err := r.prepareCommand(data)
	if err != nil {
		return nil, errors.Wrap(err, "registrarRepository SendCommand: r.prepareCommand")
	}

	log.Println("XML Request: \n", string(encoded))

	byteResponse, err := r.eppClient.Send(encoded)
	if err != nil {
		return nil, errors.Wrap(err, "registrarRepository sendXMLTCPRequest: r.eppClient.Send")
	}

	log.Println("XML Response: \n", string(byteResponse))

	return byteResponse, nil
}
