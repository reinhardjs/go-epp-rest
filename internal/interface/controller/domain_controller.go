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
	SecDNSUpdate(c *gin.Context)
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

func (controller *domainController) SecDNSUpdate(c *gin.Context) {

	AddDSDataList := []types.DSData{}
	RemoveDSDataList := []types.DSData{}

	domain := c.Query("domain")
	ext := c.Query("ext")
	isRemoveAll := c.Query("isremoveall")
	ddKeytag0 := c.Query("dd_keytag0")
	ddAlgorithm0 := c.Query("dd_algorithm0")
	ddDigestType0 := c.Query("dd_digesttype0")
	ddDigest0 := c.Query("dd_digest0")
	kdFlag0 := c.Query("kd_flag0")
	kdProtocol0 := c.Query("kd_protocol0")
	kdAlgorithm0 := c.Query("kd_algorithm0")
	kdPublicKey0 := c.Query("kd_publickey0")

	ddKeytag1 := c.Query("dd_keytag1")
	ddAlgorithm1 := c.Query("dd_algorithm1")
	ddDigestType1 := c.Query("dd_digesttype1")
	ddDigest1 := c.Query("dd_digest1")
	kdFlag1 := c.Query("kd_flag1")
	kdProtocol1 := c.Query("kd_protocol1")
	kdAlgorithm1 := c.Query("kd_algorithm1")
	kdPublicKey1 := c.Query("kd_publickey1")

	xddKeytag0 := c.Query("xdd_keytag0")
	xddAlgorithm0 := c.Query("xdd_algorithm0")
	xddDigest0 := c.Query("xdd_digest0")
	xddDigestType0 := c.Query("xdd_digesttype0")
	xkdFlag0 := c.Query("xkd_flag0")
	xkdProtocol0 := c.Query("xkd_protocol0")
	xkdAlgorithm0 := c.Query("xkd_algorithm0")
	xkdPublicKey0 := c.Query("xkd_publickey0")

	xddKeytag1 := c.Query("xdd_keytag1")
	xddAlgorithm1 := c.Query("xdd_algorithm1")
	xddDigest1 := c.Query("xdd_digest1")
	xddDigestType1 := c.Query("xdd_digesttype1")
	xkdFlag1 := c.Query("xkd_flag1")
	xkdProtocol1 := c.Query("xkd_protocol1")
	xkdAlgorithm1 := c.Query("xkd_algorithm1")
	xkdPublicKey1 := c.Query("xkd_publickey1")

	if len(strings.TrimSpace(ddKeytag0)) != 0 {
		dsData := types.DSData{
			KeyTag:     ddKeytag0,
			Alg:        ddAlgorithm0,
			DigestType: ddDigestType0,
			Digest:     ddDigest0,
		}

		if len(strings.TrimSpace(kdFlag0)) != 0 {
			dsData.KeyData = &types.KeyData{
				Flags:    kdFlag0,
				Protocol: kdProtocol0,
				Alg:      kdAlgorithm0,
				PubKey:   kdPublicKey0,
			}
		}

		AddDSDataList = append(AddDSDataList, dsData)
	}

	if len(strings.TrimSpace(ddKeytag1)) != 0 {
		dsData := types.DSData{
			KeyTag:     ddKeytag1,
			Alg:        ddAlgorithm1,
			DigestType: ddDigestType1,
			Digest:     ddDigest1,
		}

		if len(strings.TrimSpace(kdFlag1)) != 0 {
			dsData.KeyData = &types.KeyData{
				Flags:    kdFlag1,
				Protocol: kdProtocol1,
				Alg:      kdAlgorithm1,
				PubKey:   kdPublicKey1,
			}
		}

		AddDSDataList = append(AddDSDataList, dsData)
	}

	if len(strings.TrimSpace(xddKeytag0)) != 0 {
		xdsData := types.DSData{
			KeyTag:     xddKeytag0,
			Alg:        xddAlgorithm0,
			DigestType: xddDigestType0,
			Digest:     xddDigest0,
		}

		if len(strings.TrimSpace(xkdFlag0)) != 0 {
			xdsData.KeyData = &types.KeyData{
				Flags:    xkdFlag0,
				Protocol: xkdProtocol0,
				Alg:      xkdAlgorithm0,
				PubKey:   xkdPublicKey0,
			}
		}

		RemoveDSDataList = append(RemoveDSDataList, xdsData)
	}

	if len(strings.TrimSpace(xddKeytag1)) != 0 {
		xdsData := types.DSData{
			KeyTag:     xddKeytag1,
			Alg:        xddAlgorithm1,
			DigestType: xddDigestType1,
			Digest:     xddDigest1,
		}

		if len(strings.TrimSpace(xkdFlag1)) != 0 {
			xdsData.KeyData = &types.KeyData{
				Flags:    xkdFlag1,
				Protocol: xkdProtocol1,
				Alg:      xkdAlgorithm1,
				PubKey:   xkdPublicKey1,
			}
		}

		RemoveDSDataList = append(RemoveDSDataList, xdsData)
	}

	var data types.DomainUpdateType = types.DomainUpdateType{
		Command: types.DomainCommand{
			Update: types.DomainUpdate{
				Name: domain,
			},
		},
	}

	if isRemoveAll == "yes" {
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

	responseString, err := controller.interactor.SecDNSUpdate(data, ext, "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController SecDNSUpdate: controller.interactor.SecDNSUpdate"))
	}

	c.String(200, responseString)
}
