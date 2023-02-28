package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/queries"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type transferController struct {
	interactor interactor.TransferInteractor
}

type TransferController interface {
	Check(c *gin.Context)
	Request(c *gin.Context)
	Cancel(c *gin.Context)
	Approve(c *gin.Context)
	Reject(c *gin.Context)
}

func NewTransferController(interactor interactor.TransferInteractor) TransferController {
	return &transferController{
		interactor: interactor,
	}
}

func (controller *transferController) Check(c *gin.Context) {

	var transferCheckQuery queries.TransferCheckQuery
	c.ShouldBindQuery(&transferCheckQuery)

	data := types.TransferType{
		TransferParent: types.Transfer{
			Operation: "query",
			Detail: types.TransferDetail{
				Name: transferCheckQuery.Domain,
			},
		},
	}

	responseString, err := controller.interactor.Check(data, transferCheckQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "TransferController Check: controller.interactor.Check"))
	}

	c.String(200, responseString)
}

func (controller *transferController) Request(c *gin.Context) {

	var transferRequestQuery queries.TransferRequestQuery
	c.ShouldBindQuery(&transferRequestQuery)

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

	responseString, err := controller.interactor.Request(data, transferRequestQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "TransferController Request: controller.interactor.Request"))
	}

	c.String(200, responseString)
}

func (controller *transferController) Cancel(c *gin.Context) {

	var transferCancelQuery queries.TransferCancelQuery
	c.ShouldBindQuery(&transferCancelQuery)

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

	responseString, err := controller.interactor.Cancel(data, transferCancelQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "TransferController Cancel: controller.interactor.Cancel"))
	}

	c.String(200, responseString)
}

func (controller *transferController) Approve(c *gin.Context) {

	var transferApproveQuery queries.TransferApproveQuery
	c.ShouldBindQuery(&transferApproveQuery)

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

	responseString, err := controller.interactor.Approve(data, transferApproveQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "TransferController Approve: controller.interactor.Approve"))
	}

	c.String(200, responseString)
}

func (controller *transferController) Reject(c *gin.Context) {

	var transferRejectQuery queries.TransferRejectQuery
	c.ShouldBindQuery(&transferRejectQuery)

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

	responseString, err := controller.interactor.Reject(data, transferRejectQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "TransferController Reject: controller.interactor.Reject"))
	}

	c.String(200, responseString)
}
