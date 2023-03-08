package controllers

import (
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers/infrastructure"
	presenter_infrastructure "gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
)

type pollController struct {
	interactor usecase.PollInteractor
}

type PollController interface {
	Poll(c infrastructure.Context)
}

func NewPollController(interactor usecase.PollInteractor) PollController {
	return &pollController{
		interactor: interactor,
	}
}

func (controller *pollController) Poll(ctx infrastructure.Context) {
	controller.interactor.Poll(ctx.(presenter_infrastructure.Context))
}
