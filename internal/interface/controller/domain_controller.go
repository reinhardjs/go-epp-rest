package controller

import (
	"encoding/xml"
	"log"

	"github.com/bombsimon/epp-go/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

type domainController struct {
	registrarInteractor interactor.RegistrarInteractor
}

type DomainController interface {
	CheckDomain(c *gin.Context)
}

func NewDomainController(interactor interactor.RegistrarInteractor) DomainController {
	return &domainController{
		registrarInteractor: interactor,
	}
}

func (controller *domainController) CheckDomain(c *gin.Context) {
	data := types.DomainCheckType{
		Check: types.DomainCheck{
			Names: []string{"reinhard.com", "jonathan.com"},
		},
	}

	response, _ := controller.registrarInteractor.Check(data, "com", "eng")

	responseObj := model.DomainCheckResponse{}

	if err := xml.Unmarshal([]byte(response), &responseObj); err != nil {
		log.Println(errors.Wrap(err, "Domain Controller: CheckDomain xml.Unmarshal"))
	}

	c.JSON(200, gin.H{
		"message": "nine",
	})
}
