package controller

import (
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/queries"
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
	SecDNSUpdate(c *gin.Context)
}

func NewDomainController(interactor interactor.DomainInteractor) DomainController {
	return &domainController{
		interactor: interactor,
	}
}

func (controller *domainController) Check(c *gin.Context) {

	var domainCheckQuery queries.DomainCheckQuery
	c.ShouldBindQuery(&domainCheckQuery)

	domainList := strings.Split(domainCheckQuery.DomainList, ",")

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

	var domainCreateQuery queries.DomainCreateQuery
	c.ShouldBindQuery(&domainCreateQuery)

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

	responseString, err := controller.interactor.Create(data, domainCreateQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController Create: controller.interactor.Create"))
	}

	c.String(200, responseString)
}

func (controller *domainController) Delete(c *gin.Context) {

	var domainDeleteQuery queries.DomainDeleteQuery
	c.ShouldBindQuery(&domainDeleteQuery)

	data := types.DomainDeleteType{
		Delete: types.DomainDelete{
			Name: domainDeleteQuery.Domain,
		},
	}

	responseString, err := controller.interactor.Delete(data, domainDeleteQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController Delete: controller.interactor.Delete"))
	}

	c.String(200, responseString)
}

func (controller *domainController) Info(c *gin.Context) {

	var domainInfoQuery queries.DomainInfoQuery
	c.ShouldBindQuery(&domainInfoQuery)

	data := types.DomainInfoType{
		Info: types.DomainInfo{
			Name: types.DomainInfoName{
				Name: domainInfoQuery.Domain,
			},
		},
	}

	responseString, err := controller.interactor.Info(data, domainInfoQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController Info: controller.interactor.Info"))
	}

	c.String(200, responseString)
}

func (controller *domainController) SecDNSUpdate(c *gin.Context) {

	AddDSDataList := []types.DSData{}
	RemoveDSDataList := []types.DSData{}

	var secDNSUpdateQuery queries.SecDNSUpdateQuery

	c.ShouldBindQuery(&secDNSUpdateQuery)

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

	responseString, err := controller.interactor.SecDNSUpdate(data, secDNSUpdateQuery.Extension, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController SecDNSUpdate: controller.interactor.SecDNSUpdate"))
	}

	c.String(200, responseString)
}
