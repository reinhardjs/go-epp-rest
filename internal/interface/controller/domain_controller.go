package controller

import (
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type domainController struct {
	interactor interactor.DomainInteractor
}

type DomainController interface {
	Check(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Info(c *gin.Context)
}

func NewDomainController(interactor interactor.DomainInteractor) DomainController {
	return &domainController{
		interactor: interactor,
	}
}

func (controller *domainController) Check(c *gin.Context) {

	domainList := strings.Split(c.Query("domainlist"), ",")

	data := types.DomainCheckType{
		Check: types.DomainCheck{
			Names: domainList,
		},
	}

	responseString, err := controller.interactor.Check(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController Check: controller.interactor.Check"))
	}

	c.String(200, responseString)
}

func (controller *domainController) Create(c *gin.Context) {

	domain := c.Query("domain")
	ns := strings.Split(c.Query("ns"), ",")
	registrantContact := c.Query("regcon")
	adminContact := c.Query("admcon")
	techContact := c.Query("techcon")
	billingContact := c.Query("bilcon")
	authInfo := c.Query("authinfo")
	period, err := strconv.Atoi(c.Query("period"))
	ext := c.Query("ext")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController Create"))
	}

	data := types.DomainCreateType{
		Create: types.DomainCreate{
			Name: domain,
			Period: types.Period{
				Value: period,
				Unit:  "y", // yearly
			},
			NameServer: &types.NameServer{
				HostObject: ns,
			},
			Registrant: registrantContact,
			Contacts: []types.Contact{
				{
					Name: adminContact,
					Type: "admin",
				},
				{
					Name: techContact,
					Type: "tech",
				},
				{
					Name: billingContact,
					Type: "billing",
				},
			},
			AuthInfo: &types.AuthInfo{
				Password: authInfo,
			},
		},
	}

	responseString, err := controller.interactor.Create(data, ext, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController Create: controller.interactor.Create"))
	}

	c.String(200, responseString)
}

func (controller *domainController) Delete(c *gin.Context) {

	domain := c.Query("domain")
	ext := c.Query("ext")

	data := types.DomainDeleteType{
		Delete: types.DomainDelete{
			Name: domain,
		},
	}

	responseString, err := controller.interactor.Delete(data, ext, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController Delete: controller.interactor.Delete"))
	}

	c.String(200, responseString)
}

func (controller *domainController) Info(c *gin.Context) {

	domain := c.Query("domain")
	ext := c.Query("ext")

	data := types.DomainInfoType{
		Info: types.DomainInfo{
			Name: types.DomainInfoName{
				Name: domain,
			},
		},
	}

	responseString, err := controller.interactor.Info(data, ext, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController Info: controller.interactor.Info"))
	}

	c.String(200, responseString)
}
