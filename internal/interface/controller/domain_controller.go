package controller

import (
	"encoding/xml"
	"log"

	"github.com/bombsimon/epp-go/types"
	"github.com/gin-gonic/gin"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

type domainController struct {
	registrarInteractor interactor.RegistrarInteractor
}

type DomainController interface {
	CheckDomain(c *gin.Context) error
}

func NewDomainController(interactor interactor.RegistrarInteractor) DomainController {
	return &domainController{
		registrarInteractor: interactor,
	}
}

func (controller *domainController) CheckDomain(c *gin.Context) error {
	data := types.DomainCheckType{
		Check: types.DomainCheck{
			Names: []string{"reinhard.com", "jonathan.com"},
		},
	}

	response, _ := controller.registrarInteractor.Check(data, "com", "eng")

	log.Println("Response :", string(response))

	responseObj := model.DomainCheckResponse{}

	if err := xml.Unmarshal([]byte(response), &responseObj); err != nil {
		log.Println("Error :", err)
	}

	log.Println("Response Obj :", responseObj)
	return nil
}
