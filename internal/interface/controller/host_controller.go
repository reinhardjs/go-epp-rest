package controller

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/queries"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type hostController struct {
	interactor interactor.HostInteractor
}

type HostController interface {
	Check(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Info(c *gin.Context)
}

func NewHostController(interactor interactor.HostInteractor) HostController {
	return &hostController{
		interactor: interactor,
	}
}

func (controller *hostController) Check(c *gin.Context) {

	var hostCheckQuery queries.HostCheckQuery
	c.ShouldBindQuery(&hostCheckQuery)

	hostList := strings.Split(hostCheckQuery.HostList, ",")

	data := types.HostCheckType{
		Check: types.HostCheck{
			Names: hostList,
		},
	}

	responseString, err := controller.interactor.Check(data, hostCheckQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "HostController Check: controller.interactor.Check"))
	}

	c.String(200, responseString)
}

func (controller *hostController) Create(c *gin.Context) {

	var hostCreateQuery queries.HostCreateQuery
	c.ShouldBindQuery(&hostCreateQuery)

	hostName := hostCreateQuery.DNSList

	if hostName == "" {
		hostName = hostCreateQuery.Host
	}

	ipAddressList := strings.Split(hostCreateQuery.IPList, ",")
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

	responseString, err := controller.interactor.Create(data, hostCreateQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "HostController Create: controller.interactor.Create"))
	}

	c.String(200, responseString)
}

func (controller *hostController) Update(c *gin.Context) {

	var hostUpdateQuery queries.HostUpdateQuery
	c.ShouldBindQuery(&hostUpdateQuery)

	hostName := hostUpdateQuery.DNSList

	if hostName == "" {
		hostName = hostUpdateQuery.Host
	}

	addIPAddressList := strings.Split(hostUpdateQuery.AddIPList, ",")
	addHostAddressList := []types.HostAddress{}

	removeIPAddressList := strings.Split(hostUpdateQuery.RemoveIPList, ",")
	removeHostAddressList := []types.HostAddress{}

	for _, ipAddress := range addIPAddressList {
		ipType := types.HostIPv4 // need to check ip type based on ip address
		addHostAddressList = append(addHostAddressList, types.HostAddress{
			Address: ipAddress,
			IP:      ipType,
		})
	}

	for _, ipAddress := range removeIPAddressList {
		ipType := types.HostIPv4 // need to check ip type based on ip address
		removeHostAddressList = append(removeHostAddressList, types.HostAddress{
			Address: ipAddress,
			IP:      ipType,
		})
	}

	var add types.HostAddRemove
	var rem types.HostAddRemove

	if len(addHostAddressList) > 0 {
		add = types.HostAddRemove{
			Address: addHostAddressList,
		}
	}

	if len(removeHostAddressList) > 0 {
		rem = types.HostAddRemove{
			Address: removeHostAddressList,
		}
	}

	data := types.HostUpdateType{
		Update: types.HostUpdate{
			Name:   hostName,
			Add:    &add,
			Remove: &rem,
		},
	}

	responseString, err := controller.interactor.Update(data, hostUpdateQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "HostController Update: controller.interactor.Update"))
	}

	c.String(200, responseString)
}

func (controller *hostController) Delete(c *gin.Context) {

	var hostDeleteQuery queries.HostDeleteQuery
	c.ShouldBindQuery(&hostDeleteQuery)

	hostName := hostDeleteQuery.DNSList

	if hostName == "" {
		hostName = hostDeleteQuery.Host
	}

	data := types.HostDeleteType{
		Delete: types.HostDelete{
			Name: hostName,
		},
	}

	responseString, err := controller.interactor.Delete(data, hostDeleteQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "HostController Delete: controller.interactor.Delete"))
	}

	c.String(200, responseString)
}

func (controller *hostController) Info(c *gin.Context) {

	var hostInfoQuery queries.HostInfoQuery
	c.ShouldBindQuery(&hostInfoQuery)

	hostName := hostInfoQuery.DNSList

	if hostName == "" {
		hostName = hostInfoQuery.Host
	}

	data := types.HostInfoType{
		Info: types.HostInfo{
			Name: hostName,
		},
	}

	responseString, err := controller.interactor.Info(data, hostInfoQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "HostController Info: controller.interactor.Info"))
	}

	c.String(200, responseString)
}
