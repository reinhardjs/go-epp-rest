package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
)

type pollController struct {
	interactor usecase.PollInteractor
}

type PollController interface {
	Poll(c *gin.Context)
}

func NewPollController(interactor usecase.PollInteractor) PollController {
	return &pollController{
		interactor: interactor,
	}
}

func (controller *pollController) Poll(ctx *gin.Context) {
	err := controller.interactor.Poll(ctx)
	if err != nil {
		err = errors.Wrap(err, "PollController Poll")
		ctx.Error(err)
	}
}
