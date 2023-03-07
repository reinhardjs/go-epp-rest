package controllers

import (
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/request"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type domainController struct {
	interactor usecase.DomainInteractor
}

type DomainController interface {
	Check(c infrastructure.Context)
	Create(c infrastructure.Context)
	Delete(c infrastructure.Context)
	Info(c infrastructure.Context)
	SecDNSUpdate(c infrastructure.Context)
	ContactUpdate(c infrastructure.Context)
}

func NewDomainController(interactor usecase.DomainInteractor) DomainController {
	return &domainController{
		interactor: interactor,
	}
}

func (controller *domainController) Check(ctx infrastructure.Context) {

	var domainCheckQuery request.DomainCheckQuery
	ctx.BindQuery(&domainCheckQuery)

	domainList := strings.Split(domainCheckQuery.DomainList, ",")

	data := types.DomainCheckType{
		Check: types.DomainCheck{
			Names: domainList,
		},
	}

	controller.interactor.Check(ctx, data, "com", "eng")
}

func (controller *domainController) Create(ctx infrastructure.Context) {

	var domainCreateQuery request.DomainCreateQuery
	ctx.BindQuery(&domainCreateQuery)

	ns := strings.Split(domainCreateQuery.Nameserver, ",")
	period, err := strconv.Atoi(domainCreateQuery.Period)

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController Create"))
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

	controller.interactor.Create(ctx, data, domainCreateQuery.Extension, "eng")
}

func (controller *domainController) Delete(ctx infrastructure.Context) {

	var domainDeleteQuery request.DomainDeleteQuery
	ctx.BindQuery(&domainDeleteQuery)

	data := types.DomainDeleteType{
		Delete: types.DomainDelete{
			Name: domainDeleteQuery.Domain,
		},
	}

	controller.interactor.Delete(ctx, data, domainDeleteQuery.Extension, "eng")
}

func (controller *domainController) Info(ctx infrastructure.Context) {

	var domainInfoQuery request.DomainInfoQuery
	ctx.BindQuery(&domainInfoQuery)

	data := types.DomainInfoType{
		Info: types.DomainInfo{
			Name: types.DomainInfoName{
				Name: domainInfoQuery.Domain,
			},
		},
	}

	controller.interactor.Info(ctx, data, domainInfoQuery.Extension, "eng")
}

func (controller *domainController) SecDNSUpdate(ctx infrastructure.Context) {

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

	if secDNSUpdateQuery.IsRemoveAll == "yes" {
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

	controller.interactor.SecDNSUpdate(ctx, data, secDNSUpdateQuery.Extension, "eng")
}

func (controller *domainController) ContactUpdate(ctx infrastructure.Context) {

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

	controller.interactor.ContactUpdate(ctx, data, domainContactUpdateQuery.Extension, "eng")
}
