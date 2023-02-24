package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type transferPresenter struct{}

func NewTransferPresenter() presenter.TransferPresenter {
	return &transferPresenter{}
}

func (p *transferPresenter) MapCheckResponse(response []byte) (responseObject model.TransferCheckResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "TransferPresenter MapCheckResponse: xml.Unmarshal"))
	}

	return
}

func (p *transferPresenter) MapRequestResponse(response []byte) (responseObject model.TransferRequestResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "TransferPresenter MapRequestResponse: xml.Unmarshal"))
	}

	return
}

func (p *transferPresenter) MapCancelResponse(response []byte) (responseObject model.TransferCancelResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "TransferPresenter MapCancelResponse: xml.Unmarshal"))
	}

	return
}

func (p *transferPresenter) MapApproveResponse(response []byte) (responseObject model.TransferApproveResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "TransferPresenter MapApproveResponse: xml.Unmarshal"))
	}

	return
}

func (p *transferPresenter) MapRejectResponse(response []byte) (responseObject model.TransferRejectResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "TransferPresenter MapRejectResponse: xml.Unmarshal"))
	}

	return
}
