package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/common/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type transferPresenter struct{}

func NewTransferPresenter() presenter.TransferPresenter {
	return &transferPresenter{}
}

func (p *transferPresenter) Check(response []byte) (responseObject response.TransferCheckResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "TransferPresenter Check: xml.Unmarshal"))
	}

	return
}

func (p *transferPresenter) Request(response []byte) (responseObject response.TransferRequestResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "TransferPresenter Request: xml.Unmarshal"))
	}

	return
}

func (p *transferPresenter) Cancel(response []byte) (responseObject response.TransferCancelResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "TransferPresenter Cancel: xml.Unmarshal"))
	}

	return
}

func (p *transferPresenter) Approve(response []byte) (responseObject response.TransferApproveResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "TransferPresenter Approve: xml.Unmarshal"))
	}

	return
}

func (p *transferPresenter) Reject(response []byte) (responseObject response.TransferRejectResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "TransferPresenter Reject: xml.Unmarshal"))
	}

	return
}
