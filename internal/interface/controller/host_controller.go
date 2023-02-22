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

	hostName := c.Query("dnslist")
	ext := c.Query("ext")

	if hostName == "" {
		hostName = c.Query("host")
	}

	ipAddressList := strings.Split(c.Query("iplist"), ",")
	hostAddressList := []types.HostAddress{}

	for _, ipAddress := range ipAddressList {
		ipType := types.HostIPv4 // need to check ip type based on ip address
		hostAddressList = append(hostAddressList, types.HostAddress{
			Address: ipAddress,
			IP:      ipType,
		})
	}

	data := types.HostCreateType{
		Create: types.HostCreate{
			Name:    hostName,
			Address: hostAddressList,
		},
	}

	responseString, err := controller.interactor.Create(data, ext, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "HostController Create: controller.interactor.Create"))
	}

	c.String(200, responseString)
}
