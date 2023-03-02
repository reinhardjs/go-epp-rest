package controllers

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/dto/request"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type contactController struct {
	interactor interactor.ContactInteractor
}

type ContactController interface {
	Check(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Info(c *gin.Context)
}

func NewContactController(interactor interactor.ContactInteractor) ContactController {
	return &contactController{
		interactor: interactor,
	}
}

func (controller *contactController) Check(c *gin.Context) {

	var contactCheckQuery request.ContactCheckQuery
	c.ShouldBindQuery(&contactCheckQuery)

	contactList := strings.Split(contactCheckQuery.ContactList, ",")

	data := types.ContactCheckType{
		Check: types.ContactCheck{
			Names: contactList,
		},
	}

	responseString, err := controller.interactor.Check(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "ContactController Check: controller.interactor.Check"))
	}

	c.String(200, responseString)
}

func (controller *contactController) Create(c *gin.Context) {

	var contactCreateQuery request.ContactCreateQuery
	c.ShouldBindQuery(&contactCreateQuery)
	name := contactCreateQuery.Firstname + " " + contactCreateQuery.Lastname

	data := types.ContactCreateType{
		Create: types.ContactCreate{
			ID:    contactCreateQuery.Contact,
			Email: contactCreateQuery.Email,
			AuthInfo: types.AuthInfo{
				Password: contactCreateQuery.AuthInfo,
			},
			Voice: types.E164Type{
				Value: contactCreateQuery.Phone,
			},
			Fax: types.E164Type{
				Value: contactCreateQuery.Fax,
			},
			PostalInfo: []types.PostalInfo{
				{
					Name:         name,
					Organization: contactCreateQuery.Company,
					Address: types.Address{
						Street:        []string{contactCreateQuery.Address1, contactCreateQuery.Address2},
						City:          contactCreateQuery.City,
						StateProvince: contactCreateQuery.State,
						PostalCode:    contactCreateQuery.Zip,
						CountryCode:   contactCreateQuery.Country,
					},
					Type: types.PostalInfoType(types.PostalInfoInternational),
				},
			},
		},
	}

	responseString, err := controller.interactor.Create(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "ContactController Create: controller.interactor.Create"))
	}

	c.String(200, responseString)
}

func (controller *contactController) Update(c *gin.Context) {

	var contactUpdateQuery request.ContactUpdateQuery
	c.ShouldBindQuery(&contactUpdateQuery)
	// authInfo := c.Query("authinfo")
	name := contactUpdateQuery.Firstname + " " + contactUpdateQuery.Lastname

	data := types.ContactUpdateType{
		Update: types.ContactUpdate{
			Name: contactUpdateQuery.Contact,
			Change: &types.ContactChange{
				Email: contactUpdateQuery.Email,
				Voice: types.E164Type{
					Value: contactUpdateQuery.Phone,
				},
				Fax: types.E164Type{
					Value: contactUpdateQuery.Fax,
				},
				PostalInfo: []types.PostalInfo{
					{
						Name:         name,
						Organization: contactUpdateQuery.Company,
						Address: types.Address{
							Street:        []string{contactUpdateQuery.Address1, contactUpdateQuery.Address2},
							City:          contactUpdateQuery.City,
							StateProvince: contactUpdateQuery.State,
							PostalCode:    contactUpdateQuery.Zip,
							CountryCode:   contactUpdateQuery.Country,
						},
						Type: types.PostalInfoType(types.PostalInfoInternational),
					},
				},
			},
		},
	}

	responseString, err := controller.interactor.Create(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "ContactController Update: controller.interactor.Update"))
	}

	c.String(200, responseString)
}

func (controller *contactController) Delete(c *gin.Context) {

	var contactDeleteQuery request.ContactDeleteQuery
	c.ShouldBindQuery(&contactDeleteQuery)

	data := types.ContactDeleteType{
		Delete: types.ContactDelete{
			Name: contactDeleteQuery.Contact,
		},
	}

	responseString, err := controller.interactor.Delete(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "ContactController Delete: controller.interactor.Delete"))
	}

	c.String(200, responseString)
}

func (controller *contactController) Info(c *gin.Context) {

	var contactInfoQuery request.ContactInfoQuery
	c.ShouldBindQuery(&contactInfoQuery)

	data := types.ContactInfoType{
		Info: types.ContactInfo{
			Name: contactInfoQuery.Contact,
		},
	}

	responseString, err := controller.interactor.Info(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "ContactController Info: controller.interactor.Info"))
	}

	c.String(200, responseString)
}
