package interactor

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter/mapper"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type transferInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.TransferPresenter
	xmlMapper           mapper.XMLMapper
}

func NewTransferInteractor(repository repository.RegistrarRepository, presenter presenter.TransferPresenter, xmlMapper mapper.XMLMapper) usecase.TransferInteractor {
	return &transferInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
		xmlMapper:           xmlMapper,
	}
}

func (interactor *transferInteractor) Check(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Check: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.TransferCheckResponse{}
	err = interactor.xmlMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	res = interactor.Presenter.Check(*responseDTO)
	return
}

func (interactor *transferInteractor) Request(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.TransferRequestResponse{}
	err = interactor.xmlMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	res = interactor.Presenter.Request(*responseDTO)
	return
}

func (interactor *transferInteractor) Cancel(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.TransferCancelResponse{}
	err = interactor.xmlMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	res = interactor.Presenter.Cancel(*responseDTO)
	return
}

func (interactor *transferInteractor) Approve(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.TransferApproveResponse{}
	err = interactor.xmlMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	res = interactor.Presenter.Approve(*responseDTO)
	return
}

func (interactor *transferInteractor) Reject(data interface{}, ext string, langTag string) (res string, err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.TransferRejectResponse{}
	err = interactor.xmlMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	res = interactor.Presenter.Reject(*responseDTO)
	return
}
