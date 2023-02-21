package controller

import (
	"log"
	"strings"

	"github.com/bombsimon/epp-go/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

type contactController struct {
	registrarInteractor interactor.ContactInteractor
}

type ContactController interface {
	CheckContact(c *gin.Context)
}

func NewContactController(interactor interactor.ContactInteractor) ContactController {
	return &contactController{
		registrarInteractor: interactor,
	}
}

func (controller *contactController) CheckContact(c *gin.Context) {

	contactList := strings.Split(c.Query("contactlist"), ",")

	data := types.ContactCheckType{
		Check: types.ContactCheck{
			Names: contactList,
		},
	}

	responseString, err := controller.registrarInteractor.Check(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "ContactController CheckContact: controller.registrarInteractor.Check"))
	}

	c.String(200, responseString)
}
