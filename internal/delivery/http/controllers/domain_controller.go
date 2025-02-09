package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/request"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
	"golang.org/x/net/idna"
)

type domainController struct {
	interactor usecase.DomainInteractor
}

type DomainController interface {
	Check(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Info(c *gin.Context)
	SecDNSUpdate(c *gin.Context)
	ContactUpdate(c *gin.Context)
	StatusUpdate(c *gin.Context)
	AuthInfoUpdate(c *gin.Context)
	NameserverUpdate(c *gin.Context)
	Renew(c *gin.Context)
}

func NewDomainController(interactor usecase.DomainInteractor) DomainController {
	return &domainController{
		interactor: interactor,
	}
}

func (controller *domainController) Check(ctx *gin.Context) {
	var domainCheckQuery request.DomainCheckQuery
	ctx.BindQuery(&domainCheckQuery)

	domainList := strings.Split(domainCheckQuery.DomainList, ",")

	data := types.DomainCheckType{
		Check: types.DomainCheck{
			Names: domainList,
		},
	}

	err := controller.interactor.Check(ctx, data, "com", "eng")
	if err != nil {
		err = errors.Wrap(err, "DomainController Check")
		ctx.Error(err)
	}
}

func (controller *domainController) Create(ctx *gin.Context) {
	var domainCreateQuery request.DomainCreateQuery
	ctx.BindQuery(&domainCreateQuery)

	ns := []string{}
	for i := 1; i <= 13; i++ {
		nameServer := ctx.Query(fmt.Sprintf("ns%v", i))

		if nameServer != "" {
			ns = append(ns, nameServer)
		}
	}

	period, err := strconv.Atoi(domainCreateQuery.Period)
	if err != nil {
		err = errors.Wrap(err, "DomainController Check: strconv.Atoi")
		ctx.Error(err)
	}

	data := types.DomainCreateType{
		Create: types.DomainCreate{
			Name: domainCreateQuery.Domain,
			Period: types.Period{
				Value: period,
				Unit:  "y", // yearly
			},
			NameServer: &types.NameServer{
				HostObject: ns,
			},
			Registrant: domainCreateQuery.RegistrantContact,
			Contacts: []types.Contact{
				{
					Name: domainCreateQuery.AdminContact,
					Type: "admin",
				},
				{
					Name: domainCreateQuery.TechContact,
					Type: "tech",
				},
				{
					Name: domainCreateQuery.BillingContact,
					Type: "billing",
				},
			},
			AuthInfo: &types.AuthInfo{
				Password: domainCreateQuery.AuthInfo,
			},
		},
	}

	err = controller.interactor.Create(ctx, data, domainCreateQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "DomainController Create")
		ctx.Error(err)
	}
}

func (controller *domainController) Delete(ctx *gin.Context) {
	var domainDeleteQuery request.DomainDeleteQuery
	ctx.BindQuery(&domainDeleteQuery)

	data := types.DomainDeleteType{
		Delete: types.DomainDelete{
			Name: domainDeleteQuery.Domain,
		},
	}

	err := controller.interactor.Delete(ctx, data, domainDeleteQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "DomainController Delete")
		ctx.Error(err)
	}
}

func (controller *domainController) Info(ctx *gin.Context) {
	var domainInfoQuery request.DomainInfoQuery
	ctx.BindQuery(&domainInfoQuery)

	domain := domainInfoQuery.Domain
	domain, err := idna.ToASCII(domain)
	if err != nil {
		err = errors.Wrap(err, "DomainController Info: idna.ToASCII")
		return
	}

	data := types.DomainInfoType{
		Info: types.DomainInfo{
			Name: types.DomainInfoName{
				Name: domain,
			},
		},
	}

	err = controller.interactor.Info(ctx, data, domainInfoQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "DomainController Info")
		ctx.Error(err)
	}
}

func (controller *domainController) SecDNSUpdate(ctx *gin.Context) {
	AddDSDataList := []types.DSData{}
	RemoveDSDataList := []types.DSData{}

	var secDNSUpdateQuery request.SecDNSUpdateQuery

	ctx.BindQuery(&secDNSUpdateQuery)

	if len(strings.TrimSpace(secDNSUpdateQuery.DdKeytag0)) != 0 {
		dsData := types.DSData{
			KeyTag:     secDNSUpdateQuery.DdKeytag0,
			Alg:        secDNSUpdateQuery.DdAlgorithm0,
			DigestType: secDNSUpdateQuery.DdDigestType0,
			Digest:     secDNSUpdateQuery.DdDigest0,
		}

		if len(strings.TrimSpace(secDNSUpdateQuery.KdFlag0)) != 0 {
			dsData.KeyData = &types.KeyData{
				Flags:    secDNSUpdateQuery.KdFlag0,
				Protocol: secDNSUpdateQuery.KdProtocol0,
				Alg:      secDNSUpdateQuery.KdAlgorithm0,
				PubKey:   secDNSUpdateQuery.KdPublicKey0,
			}
		}

		AddDSDataList = append(AddDSDataList, dsData)
	}

	if len(strings.TrimSpace(secDNSUpdateQuery.DdKeytag1)) != 0 {
		dsData := types.DSData{
			KeyTag:     secDNSUpdateQuery.DdKeytag1,
			Alg:        secDNSUpdateQuery.DdAlgorithm1,
			DigestType: secDNSUpdateQuery.DdDigestType1,
			Digest:     secDNSUpdateQuery.DdDigest1,
		}

		if len(strings.TrimSpace(secDNSUpdateQuery.KdFlag1)) != 0 {
			dsData.KeyData = &types.KeyData{
				Flags:    secDNSUpdateQuery.KdFlag1,
				Protocol: secDNSUpdateQuery.KdProtocol1,
				Alg:      secDNSUpdateQuery.KdAlgorithm1,
				PubKey:   secDNSUpdateQuery.KdPublicKey1,
			}
		}

		AddDSDataList = append(AddDSDataList, dsData)
	}

	if len(strings.TrimSpace(secDNSUpdateQuery.XddKeytag0)) != 0 {
		xdsData := types.DSData{
			KeyTag:     secDNSUpdateQuery.XddKeytag0,
			Alg:        secDNSUpdateQuery.XddAlgorithm0,
			DigestType: secDNSUpdateQuery.XddDigestType0,
			Digest:     secDNSUpdateQuery.XddDigest0,
		}

		if len(strings.TrimSpace(secDNSUpdateQuery.XkdFlag0)) != 0 {
			xdsData.KeyData = &types.KeyData{
				Flags:    secDNSUpdateQuery.XkdFlag0,
				Protocol: secDNSUpdateQuery.XkdProtocol0,
				Alg:      secDNSUpdateQuery.XkdAlgorithm0,
				PubKey:   secDNSUpdateQuery.XkdPublicKey0,
			}
		}

		RemoveDSDataList = append(RemoveDSDataList, xdsData)
	}

	if len(strings.TrimSpace(secDNSUpdateQuery.XddKeytag1)) != 0 {
		xdsData := types.DSData{
			KeyTag:     secDNSUpdateQuery.XddKeytag1,
			Alg:        secDNSUpdateQuery.XddAlgorithm1,
			DigestType: secDNSUpdateQuery.XddDigestType1,
			Digest:     secDNSUpdateQuery.XddDigest1,
		}

		if len(strings.TrimSpace(secDNSUpdateQuery.XkdFlag1)) != 0 {
			xdsData.KeyData = &types.KeyData{
				Flags:    secDNSUpdateQuery.XkdFlag1,
				Protocol: secDNSUpdateQuery.XkdProtocol1,
				Alg:      secDNSUpdateQuery.XkdAlgorithm1,
				PubKey:   secDNSUpdateQuery.XkdPublicKey1,
			}
		}

		RemoveDSDataList = append(RemoveDSDataList, xdsData)
	}

	var data types.DomainUpdateType = types.DomainUpdateType{
		Command: types.DomainCommand{
			Update: types.DomainUpdate{
				Name: secDNSUpdateQuery.Domain,
			},
		},
	}

	if secDNSUpdateQuery.IsRemoveAll == "Y" {
		doRemove := true
		data.Command.Extension = &types.Extension{
			SecDNSUpdate: &types.SecDNSUpdate{
				Remove: &types.SecDNSAddRem{
					All: &doRemove,
				},
			},
		}
	} else {
		addData := &types.SecDNSAddRem{
			DSDatas: AddDSDataList,
		}

		remData := &types.SecDNSAddRem{
			DSDatas: RemoveDSDataList,
		}

		if len(RemoveDSDataList) == 0 {
			remData = nil
		}

		data.Command.Extension = &types.Extension{
			SecDNSUpdate: &types.SecDNSUpdate{
				Add:    addData,
				Remove: remData,
			},
		}
	}

	err := controller.interactor.SecDNSUpdate(ctx, data, secDNSUpdateQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "DomainController SecDNSUpdate")
		ctx.Error(err)
	}
}

func (controller *domainController) ContactUpdate(ctx *gin.Context) {
	var domainContactUpdateQuery request.DomainContactUpdateQuery
	ctx.BindQuery(&domainContactUpdateQuery)

	var addData, remData types.DomainAddRemove
	var chgData types.DomainChange

	var addContacts, remContacts []types.Contact
	addContacts = []types.Contact{}
	remContacts = []types.Contact{}

	if domainContactUpdateQuery.AdminContact != domainContactUpdateQuery.DeleteAdminContact {
		addContacts = append(addContacts, types.Contact{
			Name: domainContactUpdateQuery.AdminContact,
			Type: "admin",
		})
		remContacts = append(remContacts, types.Contact{
			Name: domainContactUpdateQuery.DeleteAdminContact,
			Type: "admin",
		})
	}

	if domainContactUpdateQuery.TechContact != domainContactUpdateQuery.DeleteTechContact {
		addContacts = append(addContacts, types.Contact{
			Name: domainContactUpdateQuery.TechContact,
			Type: "tech",
		})
		remContacts = append(remContacts, types.Contact{
			Name: domainContactUpdateQuery.DeleteTechContact,
			Type: "tech",
		})
	}

	if domainContactUpdateQuery.BillingContact != domainContactUpdateQuery.DeleteBillingContact {
		addContacts = append(addContacts, types.Contact{
			Name: domainContactUpdateQuery.BillingContact,
			Type: "billing",
		})
		remContacts = append(remContacts, types.Contact{
			Name: domainContactUpdateQuery.DeleteBillingContact,
			Type: "billing",
		})
	}

	if len(addContacts) > 0 {
		addData = types.DomainAddRemove{
			Contact: addContacts,
		}
	}

	if len(remContacts) > 0 {
		remData = types.DomainAddRemove{
			Contact: remContacts,
		}
	}

	if domainContactUpdateQuery.RegistrantContact != "" {
		chgData = types.DomainChange{
			Registrant: domainContactUpdateQuery.RegistrantContact,
		}
	}

	data := types.DomainUpdateType{
		Command: types.DomainCommand{
			Update: types.DomainUpdate{
				Name:   domainContactUpdateQuery.Domain,
				Add:    &addData,
				Remove: &remData,
				Change: &chgData,
			},
		},
	}

	err := controller.interactor.ContactUpdate(ctx, data, domainContactUpdateQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "DomainController ContactUpdate")
		ctx.Error(err)
	}
}

func (controller *domainController) StatusUpdate(ctx *gin.Context) {
	var domainStatusUpdateQuery request.DomainStatusUpdateQuery
	ctx.BindQuery(&domainStatusUpdateQuery)

	var addData, remData types.DomainAddRemove

	var addStatuses, remStatuses []types.DomainStatus
	addStatuses = []types.DomainStatus{}
	remStatuses = []types.DomainStatus{}

	if domainStatusUpdateQuery.Status == "ok" {
		remStatuses = append(remStatuses, types.DomainStatus{
			DomainStatusType: types.DomainStatusClientUpdateProhibited,
		})
		remStatuses = append(remStatuses, types.DomainStatus{
			DomainStatusType: types.DomainStatusClientDeleteProhibited,
		})
		remStatuses = append(remStatuses, types.DomainStatus{
			DomainStatusType: types.DomainStatusClientTransferProhibited,
		})
	} else if domainStatusUpdateQuery.Status == "clienthold" || domainStatusUpdateQuery.Status == "hold" {
		addStatuses = append(addStatuses, types.DomainStatus{
			DomainStatusType: types.DomainStatusClientHold,
		})
	} else if domainStatusUpdateQuery.Status == "unhold" {
		remStatuses = append(remStatuses, types.DomainStatus{
			DomainStatusType: types.DomainStatusClientHold,
		})
	} else {
		addStatuses = append(addStatuses, types.DomainStatus{
			DomainStatusType: types.DomainStatusClientUpdateProhibited,
		})
		addStatuses = append(addStatuses, types.DomainStatus{
			DomainStatusType: types.DomainStatusClientDeleteProhibited,
		})
		addStatuses = append(addStatuses, types.DomainStatus{
			DomainStatusType: types.DomainStatusClientTransferProhibited,
		})
	}

	if len(addStatuses) > 0 {
		addData = types.DomainAddRemove{
			Status: addStatuses,
		}
	}

	if len(remStatuses) > 0 {
		remData = types.DomainAddRemove{
			Status: remStatuses,
		}
	}

	data := types.DomainUpdateType{
		Command: types.DomainCommand{
			Update: types.DomainUpdate{
				Name:   domainStatusUpdateQuery.Domain,
				Add:    &addData,
				Remove: &remData,
			},
		},
	}

	err := controller.interactor.StatusUpdate(ctx, data, domainStatusUpdateQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "DomainController StatusUpdate")
		ctx.Error(err)
	}
}

func (controller *domainController) AuthInfoUpdate(ctx *gin.Context) {
	var domainAuthInfoUpdateQuery request.DomainAuthInfoUpdateQuery
	ctx.BindQuery(&domainAuthInfoUpdateQuery)

	var chgData types.DomainChange = types.DomainChange{
		AuthInfo: &types.AuthInfo{
			Password: domainAuthInfoUpdateQuery.AuthInfo,
		},
	}

	data := types.DomainUpdateType{
		Command: types.DomainCommand{
			Update: types.DomainUpdate{
				Name:   domainAuthInfoUpdateQuery.Domain,
				Change: &chgData,
			},
		},
	}

	err := controller.interactor.AuthInfoUpdate(ctx, data, domainAuthInfoUpdateQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "DomainController AuthInfoUpdate")
		ctx.Error(err)
	}
}

func (controller *domainController) NameserverUpdate(ctx *gin.Context) {
	var domainNameserverUpdateQuery request.DomainNameserverUpdateQuery
	ctx.BindQuery(&domainNameserverUpdateQuery)

	var addData, remData types.DomainAddRemove
	var addNameserverWrapper, remNameserverWrapper types.NameServer
	addNameservers := []string{}
	remNameservers := []string{}

	for i := 1; i <= 13; i++ {
		ns := ctx.Query(fmt.Sprintf("ns%v", i))
		xns := ctx.Query(fmt.Sprintf("xns%v", i))

		if ns != "" {
			addNameservers = append(addNameservers, ns)
		}

		if xns != "" {
			remNameservers = append(remNameservers, ns)
		}
	}

	addNameserverWrapper = types.NameServer{
		HostObject: addNameservers,
	}

	remNameserverWrapper = types.NameServer{
		HostObject: remNameservers,
	}

	if len(addNameserverWrapper.HostObject) > 0 {
		addData = types.DomainAddRemove{
			NameServer: &addNameserverWrapper,
		}
	}

	if len(remNameserverWrapper.HostObject) > 0 {
		remData = types.DomainAddRemove{
			NameServer: &remNameserverWrapper,
		}
	}

	data := types.DomainUpdateType{
		Command: types.DomainCommand{
			Update: types.DomainUpdate{
				Name:   domainNameserverUpdateQuery.Domain,
				Add:    &addData,
				Remove: &remData,
			},
		},
	}

	err := controller.interactor.NameserverUpdate(ctx, data, domainNameserverUpdateQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "DomainController NameserverUpdate")
		ctx.Error(err)
	}
}

func (controller *domainController) Renew(ctx *gin.Context) {
	var domainRenewQuery request.DomainRenewQuery
	ctx.BindQuery(&domainRenewQuery)

	period, err := strconv.Atoi(domainRenewQuery.Period)
	if err != nil {
		err = errors.Wrap(err, "DomainController Renew: strconv.Atoi")
		ctx.Error(err)
	}

	layoutFormat := "2006-01-02T15:04:05"
	currentExpireDate, err := time.Parse(layoutFormat, fmt.Sprintf("%vT23:59:59", domainRenewQuery.CurrentExpireDate))
	if err != nil {
		err = errors.Wrap(err, "DomainController Renew: time.Parse")
		ctx.Error(err)
	}

	data := types.DomainRenewType{
		Renew: types.DomainRenew{
			Name: domainRenewQuery.Domain,
			Period: types.Period{
				Value: period,
				Unit:  "y", //yearly
			},
			ExpireDate: currentExpireDate.Format("2006-01-02"),
		},
	}

	err = controller.interactor.Renew(ctx, data, domainRenewQuery.Extension, "eng")
	if err != nil {
		err = errors.Wrap(err, "DomainController Renew")
		ctx.Error(err)
	}
}
