package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
}

func NewTransferController(interactor interactor.TransferInteractor) TransferController {
	return &transferController{
		interactor: interactor,
	}
}

func (controller *transferController) Check(c *gin.Context) {

	domain := c.Query("domain")

	data := types.TransferType{
		TransferParent: types.Transfer{
			Operation: "query",
			Detail: types.TransferDetail{
				Name: domain,
			},
		},
	}

	responseString, err := controller.interactor.Check(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "TransferController Check: controller.interactor.Check"))
	}

	c.String(200, responseString)
}

func (controller *transferController) Request(c *gin.Context) {

	domain := c.Query("domain")
	authInfo := c.Query("authinfo")
	ext := c.Query("ext")

	data := types.TransferType{
		TransferParent: types.Transfer{
			Operation: "request",
			Detail: types.TransferDetail{
				Name: domain,
				AuthInfo: &types.AuthInfo{
					Password: authInfo,
				},
			},
		},
	}

	responseString, err := controller.interactor.Request(data, ext, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "TransferController Request: controller.interactor.Request"))
	}

	c.String(200, responseString)
}

func (controller *transferController) Cancel(c *gin.Context) {

	domain := c.Query("domain")
	authInfo := c.Query("authinfo")
	ext := c.Query("ext")

	data := types.TransferType{
		TransferParent: types.Transfer{
			Operation: "cancel",
			Detail: types.TransferDetail{
				Name: domain,
				AuthInfo: &types.AuthInfo{
					Password: authInfo,
				},
			},
		},
	}

	responseString, err := controller.interactor.Cancel(data, ext, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "TransferController Cancel: controller.interactor.Cancel"))
	}

	c.String(200, responseString)
}

func (controller *transferController) Approve(c *gin.Context) {

	domain := c.Query("domain")
	authInfo := c.Query("authinfo")
	ext := c.Query("ext")

	data := types.TransferType{
		TransferParent: types.Transfer{
			Operation: "approve",
			Detail: types.TransferDetail{
				Name: domain,
				AuthInfo: &types.AuthInfo{
					Password: authInfo,
				},
			},
		},
	}

	responseString, err := controller.interactor.Approve(data, ext, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "TransferController Approve: controller.interactor.Approve"))
	}

	c.String(200, responseString)
}
