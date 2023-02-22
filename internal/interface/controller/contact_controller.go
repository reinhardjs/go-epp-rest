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

	// domain := c.Query("domain")

	data := types.ContactCreateType{
		Create: types.ContactCreate{
			ID:    "id123",
			Email: "reinhardjsilalahi@gmail.com",
			AuthInfo: types.AuthInfo{
				Password: "qwe123*&",
			},
			Voice: types.E164Type{
				Value: "+1.7035555555",
			},
			Fax: types.E164Type{
				Value: "+1.7035555556",
			},
			PostalInfo: []types.PostalInfo{
				{
					Name:         "John Doe",
					Organization: "Example Inc.",
					Address: types.Address{
						Street:        []string{"123 Example Dr.", "Suite 100"},
						City:          "Dulles",
						StateProvince: "VA",
						PostalCode:    "20166-6503",
						CountryCode:   "US",
					},
					Type: types.PostalInfoType(types.PostalInfoInternational),
				},
			},
		},
	}

	responseString, err := controller.interactor.Create(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "DomainController Create: controller.interactor.Create"))
	}

	c.String(200, responseString)
}
