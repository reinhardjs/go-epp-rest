package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
)

type TransferPresenter interface {
	CheckSuccess(ctx infrastructure.Context, obj response.TransferCheckResponse)
	RequestSuccess(ctx infrastructure.Context, obj response.TransferRequestResponse)
	CancelSuccess(ctx infrastructure.Context, obj response.TransferCancelResponse)
	ApproveSuccess(ctx infrastructure.Context, obj response.TransferApproveResponse)
	RejectSuccess(ctx infrastructure.Context, obj response.TransferRejectResponse)
}
