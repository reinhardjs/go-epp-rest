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

func (controller contactController) Check(ctx infrastructure.Context) {
	var contactCheckQuery request.ContactCheckQuery
	ctx.BindQuery(&contactCheckQuery)

	contactList := strings.Split(contactCheckQuery.ContactList, ",")

	data := types.ContactCheckType{
		Check: types.ContactCheck{
			Names: contactList,
		},
	}

	err := controller.interactor.Check(ctx.(presenter_infrastructure.Context), data, "com", "eng")
	if err != nil {
		err = errors.Wrap(err, "ContactController Check")
		ctx.AbortWithError(200, err)
	}
}

func (controller contactController) Create(ctx infrastructure.Context) {
	var contactCreateQuery request.ContactCreateQuery
	ctx.BindQuery(&contactCreateQuery)
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

	err := controller.interactor.Create(ctx.(presenter_infrastructure.Context), data, "com", "eng")
	if err != nil {
		err = errors.Wrap(err, "ContactController Create")
		ctx.AbortWithError(200, err)
	}
}

func (controller contactController) Update(ctx infrastructure.Context) {
	var contactUpdateQuery request.ContactUpdateQuery
	ctx.BindQuery(&contactUpdateQuery)
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

	err := controller.interactor.Update(ctx.(presenter_infrastructure.Context), data, "com", "eng")
	if err != nil {
		err = errors.Wrap(err, "ContactController Update")
		ctx.AbortWithError(200, err)
	}
}

func (controller contactController) Delete(ctx infrastructure.Context) {
	var contactDeleteQuery request.ContactDeleteQuery
	ctx.BindQuery(&contactDeleteQuery)

	data := types.ContactDeleteType{
		Delete: types.ContactDelete{
			Name: contactDeleteQuery.Contact,
		},
	}

	err := controller.interactor.Delete(ctx.(presenter_infrastructure.Context), data, "com", "eng")
	if err != nil {
		err = errors.Wrap(err, "ContactController Delete")
		ctx.AbortWithError(200, err)
	}
}

func (controller contactController) Info(ctx infrastructure.Context) {
	var contactInfoQuery request.ContactInfoQuery
	ctx.BindQuery(&contactInfoQuery)

	data := types.ContactInfoType{
		Info: types.ContactInfo{
			Name: contactInfoQuery.Contact,
		},
	}

	err := controller.interactor.Info(ctx.(presenter_infrastructure.Context), data, "com", "eng")
	if err != nil {
		err = errors.Wrap(err, "ContactController Info")
		ctx.AbortWithError(200, err)
	}
}
