package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/common/dto/response"
)

type TransferPresenter interface {
	Check(response []byte) (response.TransferCheckResponse, error)
	Request(response []byte) (response.TransferRequestResponse, error)
	Cancel(response []byte) (response.TransferCancelResponse, error)
	Approve(response []byte) (response.TransferApproveResponse, error)
	Reject(response []byte) (response.TransferRejectResponse, error)
}
