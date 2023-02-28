package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/responses"
)

type TransferPresenter interface {
	MapCheckResponse(response []byte) (responses.TransferCheckResponse, error)
	MapRequestResponse(response []byte) (responses.TransferRequestResponse, error)
	MapCancelResponse(response []byte) (responses.TransferCancelResponse, error)
	MapApproveResponse(response []byte) (responses.TransferApproveResponse, error)
	MapRejectResponse(response []byte) (responses.TransferRejectResponse, error)
}
