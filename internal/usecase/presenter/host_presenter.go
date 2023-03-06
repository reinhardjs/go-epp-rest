package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
)

type HostPresenter interface {
	CheckSuccess(ctx infrastructure.Context, obj response.CheckHostResponse)
	CreateSuccess(ctx infrastructure.Context, obj response.CreateHostResponse)
	UpdateSuccess(ctx infrastructure.Context, obj response.UpdateHostResponse)
	DeleteSuccess(ctx infrastructure.Context, obj response.DeleteHostResponse)
	InfoSuccess(ctx infrastructure.Context, obj response.InfoHostResponse)
}
