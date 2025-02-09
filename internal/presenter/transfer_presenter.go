package presenter

import (
	"fmt"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/utils"
)

type transferPresenter struct{}

func NewTransferPresenter() presenter.TransferPresenter {
	return &transferPresenter{}
}

func (p *transferPresenter) CheckSuccess(ctx infrastructure.Context, responseObject response.TransferCheckResponse) (err error) {
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	buffer.WriteString(fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message))

	if responseObject.Result.ExternalValue != nil {
		buffer.WriteString(fmt.Sprintf(" | %s %s", responseObject.Result.ExternalValue.Value.ReasonCode, responseObject.Result.ExternalValue.Reason))
	}

	ctx.String(200, buffer.String())
	return
}

func (p *transferPresenter) RequestSuccess(ctx infrastructure.Context, responseObject response.TransferRequestResponse) (err error) {
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	buffer.WriteString(fmt.Sprintf("%v %v", "1000", responseObject.Result.Message))

	if responseObject.Result.ExternalValue != nil {
		buffer.WriteString(fmt.Sprintf(" | %s %s", responseObject.Result.ExternalValue.Value.ReasonCode, responseObject.Result.ExternalValue.Reason))
	}

	ctx.String(200, buffer.String())
	return
}

func (p *transferPresenter) CancelSuccess(ctx infrastructure.Context, responseObject response.TransferCancelResponse) (err error) {
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	buffer.WriteString(fmt.Sprintf("%v %v", "1000", responseObject.Result.Message))

	if responseObject.Result.ExternalValue != nil {
		buffer.WriteString(fmt.Sprintf(" | %s %s", responseObject.Result.ExternalValue.Value.ReasonCode, responseObject.Result.ExternalValue.Reason))
	}

	ctx.String(200, buffer.String())
	return
}

func (p *transferPresenter) ApproveSuccess(ctx infrastructure.Context, responseObject response.TransferApproveResponse) (err error) {
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	buffer.WriteString(fmt.Sprintf("%v %v", "1000", responseObject.Result.Message))

	if responseObject.Result.ExternalValue != nil {
		buffer.WriteString(fmt.Sprintf(" | %s %s", responseObject.Result.ExternalValue.Value.ReasonCode, responseObject.Result.ExternalValue.Reason))
	}

	ctx.String(200, buffer.String())
	return
}

func (p *transferPresenter) RejectSuccess(ctx infrastructure.Context, responseObject response.TransferRejectResponse) (err error) {
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	buffer.WriteString(fmt.Sprintf("%v %v", "1000", responseObject.Result.Message))

	if responseObject.Result.ExternalValue != nil {
		buffer.WriteString(fmt.Sprintf(" | %s %s", responseObject.Result.ExternalValue.Value.ReasonCode, responseObject.Result.ExternalValue.Reason))
	}

	ctx.String(200, buffer.String())
	return
}
