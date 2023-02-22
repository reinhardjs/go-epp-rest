package controller

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type hostController struct {
	interactor interactor.HostInteractor
}

type HostController interface {
	Check(c *gin.Context)
	Create(c *gin.Context)
}

func NewHostController(interactor interactor.HostInteractor) HostController {
	return &hostController{
		interactor: interactor,
	}
}

func (controller *hostController) Check(c *gin.Context) {

	hostList := strings.Split(c.Query("hostlist"), ",")

	data := types.HostCheckType{
		Check: types.HostCheck{
			Names: hostList,
		},
	}

	responseString, err := controller.interactor.Check(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "HostController Check: controller.interactor.Check"))
	}

	c.String(200, responseString)
}

func (controller *hostController) Create(c *gin.Context) {

	// domain := c.Query("domain")

	data := types.HostCreateType{
		Create: types.HostCreate{
			Name: "ns1.example.com",
			Address: []types.HostAddress{
				{
					Address: "1.1.1.1",
					IP:      types.HostIPv4,
				},
			},
		},
	}

	responseString, err := controller.interactor.Create(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "HostController Create: controller.interactor.Create"))
	}

	c.String(200, responseString)
}
