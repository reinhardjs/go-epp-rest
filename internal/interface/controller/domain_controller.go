package controller

import (
	"log"
	"strings"

	"github.com/bombsimon/epp-go/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

type domainController struct {
	registrarInteractor interactor.DomainInteractor
}

type DomainController interface {
	CheckDomain(c *gin.Context)
}

func NewDomainController(interactor interactor.DomainInteractor) DomainController {
	return &domainController{
		registrarInteractor: interactor,
	}
}

func (controller *domainController) CheckDomain(c *gin.Context) {

	domainList := strings.Split(c.Query("domainlist"), ",")

	data := types.DomainCheckType{
		Check: types.DomainCheck{
			Names: domainList,
		},
	}

	responseString, err := controller.registrarInteractor.Check(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController CheckDomain: controller.registrarInteractor.Check"))
	}

	c.String(200, responseString)
}
