package presenter

import (
	"fmt"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type transferPresenter struct{}

func NewTransferPresenter() presenter.TransferPresenter {
	return &transferPresenter{}
}

func (p *transferPresenter) CheckSuccess(ctx infrastructure.Context, responseObject response.TransferCheckResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
	return
}

func (p *transferPresenter) RequestSuccess(ctx infrastructure.Context, responseObject response.TransferRequestResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
	return
}

func (p *transferPresenter) CancelSuccess(ctx infrastructure.Context, responseObject response.TransferCancelResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
	return
}

func (p *transferPresenter) ApproveSuccess(ctx infrastructure.Context, responseObject response.TransferApproveResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
	return
}

func (p *transferPresenter) RejectSuccess(ctx infrastructure.Context, responseObject response.TransferRejectResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
	return
}
