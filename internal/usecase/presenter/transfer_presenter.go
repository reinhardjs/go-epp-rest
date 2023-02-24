package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
)

type TransferPresenter interface {
	MapCheckResponse(response []byte) (model.TransferCheckResponse, error)
	MapRequestResponse(response []byte) (model.TransferRequestResponse, error)
	MapCancelResponse(response []byte) (model.TransferCancelResponse, error)
	MapApproveResponse(response []byte) (model.TransferApproveResponse, error)
	MapRejectResponse(response []byte) (model.TransferRejectResponse, error)
}
