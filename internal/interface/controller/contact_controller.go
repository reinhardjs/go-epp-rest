package controller

import (
	"log"
	"strings"

	"github.com/bombsimon/epp-go/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/constraints"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

type contactController[T constraints.RegistrarResponseConstraint] struct {
	registrarInteractor interactor.RegistrarInteractor[T]
}

type ContactController interface {
	Check(c *gin.Context)
}

func NewContactController[T constraints.RegistrarResponseConstraint](interactor interactor.RegistrarInteractor[T]) ContactController {
	return &contactController[T]{
		registrarInteractor: interactor,
	}
}

func (controller *contactController[T]) Check(c *gin.Context) {

	contactList := strings.Split(c.Query("contactlist"), ",")

	data := types.ContactCheckType{
		Check: types.ContactCheck{
			Names: contactList,
		},
	}

	responseString, err := controller.registrarInteractor.Check(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "ContactController Check: controller.registrarInteractor.Check"))
	}

	c.String(200, responseString)
}
