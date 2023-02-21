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

type hostController[T constraints.RegistrarResponseConstraint] struct {
	registrarInteractor interactor.RegistrarInteractor[T]
}

type HostController interface {
	Check(c *gin.Context)
}

func NewHostController[T constraints.RegistrarResponseConstraint](interactor interactor.RegistrarInteractor[T]) HostController {
	return &hostController[T]{
		registrarInteractor: interactor,
	}
}

func (controller *hostController[T]) Check(c *gin.Context) {

	hostList := strings.Split(c.Query("hostlist"), ",")

	data := types.HostCheckType{
		Check: types.HostCheck{
			Names: hostList,
		},
	}

	responseString, err := controller.registrarInteractor.Check(data, "com", "eng")

	if err != nil {
		log.Println(errors.Wrap(err, "HostController Check: controller.registrarInteractor.Check"))
	}

	c.String(200, responseString)
}
