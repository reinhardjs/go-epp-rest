package controllers

import (
	"log"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/request"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type contactController struct {
	interactor usecase.ContactInteractor
}

type ContactController interface {
	Check(c infrastructure.Context)
	Create(c infrastructure.Context)
	Update(c infrastructure.Context)
	Delete(c infrastructure.Context)
	Info(c infrastructure.Context)
}

func NewContactController(interactor usecase.ContactInteractor) ContactController {
	return &contactController{
		interactor: interactor,
	}
}

func (controller contactController) Check(c infrastructure.Context) {

	var contactCheckQuery request.ContactCheckQuery
	c.BindQuery(&contactCheckQuery)

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

func (controller contactController) Create(c infrastructure.Context) {

	var contactCreateQuery request.ContactCreateQuery
	c.BindQuery(&contactCreateQuery)
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

func (controller contactController) Update(c infrastructure.Context) {

	var contactUpdateQuery request.ContactUpdateQuery
	c.BindQuery(&contactUpdateQuery)
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

func (controller contactController) Delete(c infrastructure.Context) {

	var contactDeleteQuery request.ContactDeleteQuery
	c.BindQuery(&contactDeleteQuery)

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

func (controller contactController) Info(c infrastructure.Context) {

	var contactInfoQuery request.ContactInfoQuery
	c.BindQuery(&contactInfoQuery)

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
