package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
)

type TransferPresenter interface {
	CheckSuccess(ctx infrastructure.Context, obj response.TransferCheckResponse) error
	RequestSuccess(ctx infrastructure.Context, obj response.TransferRequestResponse) error
	CancelSuccess(ctx infrastructure.Context, obj response.TransferCancelResponse) error
	ApproveSuccess(ctx infrastructure.Context, obj response.TransferApproveResponse) error
	RejectSuccess(ctx infrastructure.Context, obj response.TransferRejectResponse) error
}
