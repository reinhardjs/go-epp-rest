package interactor

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
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

func (interactor *transferInteractor) Check(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Check: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.TransferCheckResponse{}
	err = interactor.xmlMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Check: interactor.XMLMapper.Decode (TransferCheckResponse)")
		return
	}

	err = interactor.Presenter.CheckSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Check")
		return
	}
	return
}

func (interactor *transferInteractor) Request(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.TransferRequestResponse{}
	err = interactor.xmlMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.XMLMapper.Decode (TransferRequestResponse)")
		return
	}

	err = interactor.Presenter.RequestSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request")
		return
	}
	return
}

func (interactor *transferInteractor) Cancel(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.TransferCancelResponse{}
	err = interactor.xmlMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Cancel: interactor.XMLMapper.Decode (TransferCancelResponse)")
		return
	}

	err = interactor.Presenter.CancelSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Cancel")
		return
	}
	return
}

func (interactor *transferInteractor) Approve(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.TransferApproveResponse{}
	err = interactor.xmlMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Approve: interactor.XMLMapper.Decode (TransferApproveResponse)")
		return
	}

	err = interactor.Presenter.ApproveSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Approve")
		return
	}
	return
}

func (interactor *transferInteractor) Reject(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Request: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.TransferRejectResponse{}
	err = interactor.xmlMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Reject: interactor.XMLMapper.Decode (TransferRejectResponse)")
		return
	}

	err = interactor.Presenter.RejectSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "TransferInteractor Reject")
		return
	}
	return
}
