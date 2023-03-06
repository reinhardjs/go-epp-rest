package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
)

type ContactPresenter interface {
	CheckSuccess(ctx infrastructure.Context, obj response.CheckContactResponse)
	CreateSuccess(ctx infrastructure.Context, obj response.CreateContactResponse)
	UpdateSuccess(ctx infrastructure.Context, obj response.UpdateContactResponse)
	DeleteSuccess(ctx infrastructure.Context, obj response.DeleteContactResponse)
	InfoSuccess(ctx infrastructure.Context, obj response.InfoContactResponse)
}
