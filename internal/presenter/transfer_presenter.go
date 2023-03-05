package presenter

import (
	"fmt"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type transferPresenter struct{}

func NewTransferPresenter() presenter.TransferPresenter {
	return &transferPresenter{}
}

func (p *transferPresenter) Check(responseObject response.TransferCheckResponse) (res string) {

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	return
}

func (p *transferPresenter) Request(responseObject response.TransferRequestResponse) (res string) {

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	return
}

func (p *transferPresenter) Cancel(responseObject response.TransferCancelResponse) (res string) {

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	return
}

func (p *transferPresenter) Approve(responseObject response.TransferApproveResponse) (res string) {

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	return
}

func (p *transferPresenter) Reject(responseObject response.TransferRejectResponse) (res string) {

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	return
}
