package controllers

import (
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/request"
	presenter_infrastructure "gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type hostController struct {
	interactor usecase.HostInteractor
}

type HostController interface {
	Check(c infrastructure.Context)
	Create(c infrastructure.Context)
	Update(c infrastructure.Context)
	Delete(c infrastructure.Context)
	Info(c infrastructure.Context)
	Change(c infrastructure.Context)
	CheckAndCreate(c infrastructure.Context)
}

func NewHostController(interactor usecase.HostInteractor) HostController {
	return &hostController{
		interactor: interactor,
	}
}

func (controller *hostController) Check(ctx infrastructure.Context) {
	defer ctx.OnClose()

	var hostCheckQuery request.HostCheckQuery
	ctx.BindQuery(&hostCheckQuery)

	data := types.HostCheckType{
		Check: types.HostCheck{
			Names: []string{
				hostCheckQuery.Host,
			},
		},
	}

	err := controller.interactor.Check(ctx.(presenter_infrastructure.Context), data, hostCheckQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "HostController Check")
		ctx.AbortWithError(200, err)
	}
}

func (controller *hostController) Create(ctx infrastructure.Context) {
	defer ctx.OnClose()

	var hostCreateQuery request.HostCreateQuery
	ctx.BindQuery(&hostCreateQuery)

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
			IPType:  ipType,
		})
	}

	data := types.HostCreateType{
		Create: types.HostCreate{
			Name:    hostName,
			Address: hostAddressList,
		},
	}

	err := controller.interactor.Create(ctx.(presenter_infrastructure.Context), data, hostCreateQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "HostController Create")
		ctx.AbortWithError(200, err)
	}
}

func (controller *hostController) Update(ctx infrastructure.Context) {
	defer ctx.OnClose()

	var hostUpdateQuery request.HostUpdateQuery
	ctx.BindQuery(&hostUpdateQuery)

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
			IPType:  ipType,
		})
	}

	for _, ipAddress := range removeIPAddressList {
		ipType := types.HostIPv4 // need to check ip type based on ip address
		removeHostAddressList = append(removeHostAddressList, types.HostAddress{
			Address: ipAddress,
			IPType:  ipType,
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

	err := controller.interactor.Update(ctx.(presenter_infrastructure.Context), data, hostUpdateQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "HostController Update")
		ctx.AbortWithError(200, err)
	}
}

func (controller *hostController) Delete(ctx infrastructure.Context) {
	defer ctx.OnClose()

	var hostDeleteQuery request.HostDeleteQuery
	ctx.BindQuery(&hostDeleteQuery)

	hostName := hostDeleteQuery.DNSList

	if hostName == "" {
		hostName = hostDeleteQuery.Host
	}

	data := types.HostDeleteType{
		Delete: types.HostDelete{
			Name: hostName,
		},
	}

	err := controller.interactor.Delete(ctx.(presenter_infrastructure.Context), data, hostDeleteQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "HostController Delete")
		ctx.AbortWithError(200, err)
	}
}

func (controller *hostController) Info(ctx infrastructure.Context) {
	defer ctx.OnClose()

	var hostInfoQuery request.HostInfoQuery
	ctx.BindQuery(&hostInfoQuery)

	hostName := hostInfoQuery.DNSList

	if hostName == "" {
		hostName = hostInfoQuery.Host
	}

	data := types.HostInfoType{
		Info: types.HostInfo{
			Name: hostName,
		},
	}

	err := controller.interactor.Info(ctx.(presenter_infrastructure.Context), data, hostInfoQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "HostController Info")
		ctx.AbortWithError(200, err)
	}
}

func (controller *hostController) Change(ctx infrastructure.Context) {
	defer ctx.OnClose()

	var hostChangeQuery request.HostChangeQuery
	ctx.BindQuery(&hostChangeQuery)

	hostName := hostChangeQuery.Host

	data := types.HostUpdateType{
		Update: types.HostUpdate{
			Name: hostName,
			Add: &types.HostAddRemove{
				Address: []types.HostAddress{
					{
						Address: "190.1.1.1",
						IPType:  types.HostIPv4,
					},
				},
			},
			Remove: &types.HostAddRemove{}, // filled on hostinteractor's Change, from host info response
			Change: &types.HostChange{
				Name: hostChangeQuery.NewHost,
			},
		},
	}

	err := controller.interactor.Change(ctx.(presenter_infrastructure.Context), data, hostChangeQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "HostController Change")
		ctx.AbortWithError(200, err)
	}
}

func (controller *hostController) CheckAndCreate(ctx infrastructure.Context) {
	defer ctx.OnClose()

	var hostCreateQuery request.HostCheckAndCreateQuery
	ctx.BindQuery(&hostCreateQuery)

	hostName := hostCreateQuery.Host

	ipAddressList := strings.Split(hostCreateQuery.IPList, ",")
	hostAddressList := []types.HostAddress{}

	for _, ipAddress := range ipAddressList {
		ipType := types.HostIPv4 // need to check ip type based on ip address
		hostAddressList = append(hostAddressList, types.HostAddress{
			Address: ipAddress,
			IPType:  ipType,
		})
	}

	data := types.HostCreateType{
		Create: types.HostCreate{
			Name:    hostName,
			Address: hostAddressList,
		},
	}

	err := controller.interactor.CheckAndCreate(ctx.(presenter_infrastructure.Context), data, hostCreateQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "HostController CheckAndCreate")
		ctx.AbortWithError(200, err)
	}
}
