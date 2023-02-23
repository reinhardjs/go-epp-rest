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
