package delivery

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/interactor"
)

type pollController struct {
	interactor interactor.PollInteractor
}

type PollController interface {
	Poll(c *gin.Context)
}

func NewPollController(interactor interactor.PollInteractor) PollController {
	return &pollController{
		interactor: interactor,
	}
}

func (controller *pollController) Poll(c *gin.Context) {

	responseString, err := controller.interactor.Poll()

	if err != nil {
		log.Println(errors.Wrap(err, "PollController Poll: controller.interactor.Poll"))
	}

	c.String(200, responseString)
}
