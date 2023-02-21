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
	registrarInteractor interactor.RegistrarInteractor
}

type DomainController interface {
	Check(c *gin.Context)
}

func NewDomainController(interactor interactor.RegistrarInteractor) DomainController {
	return &domainController{
		registrarInteractor: interactor,
	}
}

func (controller *domainController) Check(c *gin.Context) {

	domainList := strings.Split(c.Query("domainlist"), ",")

	data := types.DomainCheckType{
		Check: types.DomainCheck{
			Names: domainList,
		},
	}

	responseString, err := controller.registrarInteractor.Check(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController Check: controller.registrarInteractor.Check"))
	}

	c.String(200, responseString)
}
