package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
)

type HostPresenter interface {
	Check(ctx infrastructure.Context, obj response.CheckHostResponse) error
	Create(ctx infrastructure.Context, obj response.CreateHostResponse) error
	Update(ctx infrastructure.Context, obj response.UpdateHostResponse) error
	Delete(ctx infrastructure.Context, obj response.DeleteHostResponse) error
	Info(ctx infrastructure.Context, obj response.InfoHostResponse) error
	CheckAndCreate(ctx infrastructure.Context, obj response.CreateHostResponse) error
}
