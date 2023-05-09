package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/request"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type transferController struct {
	interactor usecase.TransferInteractor
}

type TransferController interface {
	Check(c *gin.Context)
	Request(c *gin.Context)
	Cancel(c *gin.Context)
	Approve(c *gin.Context)
	Reject(c *gin.Context)
}

func NewTransferController(interactor usecase.TransferInteractor) TransferController {
	return &transferController{
		interactor: interactor,
	}
}

func (controller *transferController) Check(ctx *gin.Context) {
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

	err := controller.interactor.Check(ctx, data, transferCheckQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "TransferController Check")
		ctx.Error(err)
	}
}

func (controller *transferController) Request(ctx *gin.Context) {
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

	err := controller.interactor.Request(ctx, data, transferRequestQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "TransferController Request")
		ctx.Error(err)
	}
}

func (controller *transferController) Cancel(ctx *gin.Context) {

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

	err := controller.interactor.Cancel(ctx, data, transferCancelQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "TransferController Cancel")
		ctx.Error(err)
	}
}

func (controller *transferController) Approve(ctx *gin.Context) {
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

	err := controller.interactor.Approve(ctx, data, transferApproveQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "TransferController Approve")
		ctx.Error(err)
	}
}

func (controller *transferController) Reject(ctx *gin.Context) {
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

	err := controller.interactor.Reject(ctx, data, transferRejectQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "TransferController Reject")
		ctx.Error(err)
	}
}
