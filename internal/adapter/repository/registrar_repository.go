package repository

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/error_types"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter/mapper"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp"
)

type registrarRepository struct {
	eppClient adapter.EppClient
	XMLMapper mapper.XMLMapper
}

func NewRegistrarRepository(eppClient adapter.EppClient, xmlMapper mapper.XMLMapper) repository.RegistrarRepository {
	return &registrarRepository{eppClient, xmlMapper}
}

func (r *registrarRepository) prepareCommand(data interface{}) ([]byte, error) {
	encoded, err := registry_epp.Encode(data, registry_epp.ClientXMLAttributes())
	if err != nil {
		return nil, errors.Wrap(err, "registrarRepository prepareCommand: registry_epp.Encode")
	}

	return encoded, nil
}

func (r *registrarRepository) checkCommandError(byteResponse []byte) (err error) {
	obj := &response.Response{}
	err = r.XMLMapper.Decode(byteResponse, obj)

	if err != nil {
		err = errors.Wrap(err, "registrarRepository SendCommand: r.XMLMapper.Decode")
		return
	}

	if obj.Result.Code >= 2000 {
		err = &error_types.EPPCommandError{Result: obj.Result}
		return
	}

	return
}

func (r *registrarRepository) SendCommand(data interface{}) ([]byte, error) {
	encoded, err := r.prepareCommand(data)
	if err != nil {
		return nil, errors.Wrap(err, "registrarRepository SendCommand: r.prepareCommand")
	}

	byteResponse, err := r.eppClient.Send(encoded)
	if err != nil {
		return nil, errors.Wrap(err, "registrarRepository SendCommand: r.eppClient.Send")
	}

	err = r.checkCommandError(byteResponse)
	if err != nil {
		return nil, errors.Wrap(err, "registrarRepository SendCommand: epp command error")
	}

	return byteResponse, nil
}
