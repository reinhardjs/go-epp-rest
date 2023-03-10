package controllers

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/request"
	presenter_infrastructure "gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type transferController struct {
	interactor usecase.TransferInteractor
}

type TransferController interface {
	Check(c infrastructure.Context)
	Request(c infrastructure.Context)
	Cancel(c infrastructure.Context)
	Approve(c infrastructure.Context)
	Reject(c infrastructure.Context)
}

func NewTransferController(interactor usecase.TransferInteractor) TransferController {
	return &transferController{
		interactor: interactor,
	}
}

func (controller *transferController) Check(ctx infrastructure.Context) {

	var transferCheckQuery request.TransferCheckQuery
	ctx.BindQuery(&transferCheckQuery)

	data := types.TransferType{
		TransferParent: types.Transfer{
			Operation: "query",
			Detail: types.TransferDetail{
				Name: transferCheckQuery.Domain,
			},
		},
	}

	err := controller.interactor.Check(ctx.(presenter_infrastructure.Context), data, transferCheckQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "TransferController Check")
		ctx.AbortWithError(200, err)
	}
}

func (controller *transferController) Request(ctx infrastructure.Context) {

	var transferRequestQuery request.TransferRequestQuery
	ctx.BindQuery(&transferRequestQuery)

	data := types.TransferType{
		TransferParent: types.Transfer{
			Operation: "request",
			Detail: types.TransferDetail{
				Name: transferRequestQuery.Domain,
				AuthInfo: &types.AuthInfo{
					Password: transferRequestQuery.AuthInfo,
				},
			},
		},
	}

	err := controller.interactor.Request(ctx.(presenter_infrastructure.Context), data, transferRequestQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "TransferController Request")
		ctx.AbortWithError(200, err)
	}
}

func (controller *transferController) Cancel(ctx infrastructure.Context) {

	var transferCancelQuery request.TransferCancelQuery
	ctx.BindQuery(&transferCancelQuery)

	data := types.TransferType{
		TransferParent: types.Transfer{
			Operation: "cancel",
			Detail: types.TransferDetail{
				Name: transferCancelQuery.Domain,
				AuthInfo: &types.AuthInfo{
					Password: transferCancelQuery.AuthInfo,
				},
			},
		},
	}

	err := controller.interactor.Cancel(ctx.(presenter_infrastructure.Context), data, transferCancelQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "TransferController Cancel")
		ctx.AbortWithError(200, err)
	}
}

func (controller *transferController) Approve(ctx infrastructure.Context) {

	var transferApproveQuery request.TransferApproveQuery
	ctx.BindQuery(&transferApproveQuery)

	data := types.TransferType{
		TransferParent: types.Transfer{
			Operation: "approve",
			Detail: types.TransferDetail{
				Name: transferApproveQuery.Domain,
				AuthInfo: &types.AuthInfo{
					Password: transferApproveQuery.AuthInfo,
				},
			},
		},
	}

	err := controller.interactor.Approve(ctx.(presenter_infrastructure.Context), data, transferApproveQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "TransferController Approve")
		ctx.AbortWithError(200, err)
	}
}

func (controller *transferController) Reject(ctx infrastructure.Context) {

	var transferRejectQuery request.TransferRejectQuery
	ctx.BindQuery(&transferRejectQuery)

	data := types.TransferType{
		TransferParent: types.Transfer{
			Operation: "reject",
			Detail: types.TransferDetail{
				Name: transferRejectQuery.Domain,
				AuthInfo: &types.AuthInfo{
					Password: transferRejectQuery.AuthInfo,
				},
			},
		},
	}

	err := controller.interactor.Reject(ctx.(presenter_infrastructure.Context), data, transferRejectQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "TransferController Reject")
		ctx.AbortWithError(200, err)
	}
}
