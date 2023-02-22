package controller

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
}

func NewContactController(interactor interactor.ContactInteractor) ContactController {
	return &contactController{
		interactor: interactor,
	}
}

func (controller *contactController) Check(c *gin.Context) {

	contactList := strings.Split(c.Query("contactlist"), ",")

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

	contact := c.Query("contact")
	email := c.Query("email")
	authInfo := c.Query("authinfo")
	phone := c.Query("phone")
	fax := c.Query("fax")
	fname := c.Query("fname")
	lname := c.Query("lname")
	name := fname + " " + lname
	company := c.Query("company")
	addr1 := c.Query("addr1")
	addr2 := c.Query("addr2")
	city := c.Query("city")
	state := c.Query("state")
	zip := c.Query("zip")
	country := c.Query("country")

	data := types.ContactCreateType{
		Create: types.ContactCreate{
			ID:    contact,
			Email: email,
			AuthInfo: types.AuthInfo{
				Password: authInfo,
			},
			Voice: types.E164Type{
				Value: phone,
			},
			Fax: types.E164Type{
				Value: fax,
			},
			PostalInfo: []types.PostalInfo{
				{
					Name:         name,
					Organization: company,
					Address: types.Address{
						Street:        []string{addr1, addr2},
						City:          city,
						StateProvince: state,
						PostalCode:    zip,
						CountryCode:   country,
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

	contact := c.Query("contact")
	email := c.Query("email")
	// authInfo := c.Query("authinfo")
	phone := c.Query("phone")
	fax := c.Query("fax")
	fname := c.Query("fname")
	lname := c.Query("lname")
	name := fname + " " + lname
	company := c.Query("company")
	addr1 := c.Query("addr1")
	addr2 := c.Query("addr2")
	city := c.Query("city")
	state := c.Query("state")
	zip := c.Query("zip")
	country := c.Query("country")

	data := types.ContactUpdateType{
		Update: types.ContactUpdate{
			Name: contact,
			Change: &types.ContactChange{
				Email: email,
				Voice: types.E164Type{
					Value: phone,
				},
				Fax: types.E164Type{
					Value: fax,
				},
				PostalInfo: []types.PostalInfo{
					{
						Name:         name,
						Organization: company,
						Address: types.Address{
							Street:        []string{addr1, addr2},
							City:          city,
							StateProvince: state,
							PostalCode:    zip,
							CountryCode:   country,
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
