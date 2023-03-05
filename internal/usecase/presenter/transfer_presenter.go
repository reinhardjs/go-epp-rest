package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
)

type TransferPresenter interface {
	Check(response.TransferCheckResponse) string
	Request(response.TransferRequestResponse) string
	Cancel(response.TransferCancelResponse) string
	Approve(response.TransferApproveResponse) string
	Reject(response.TransferRejectResponse) string
}
