package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
)

type ContactPresenter interface {
	CheckSuccess(ctx infrastructure.Context, obj response.CheckContactResponse) error
	CreateSuccess(ctx infrastructure.Context, obj response.CreateContactResponse) error
	UpdateSuccess(ctx infrastructure.Context, obj response.UpdateContactResponse) error
	DeleteSuccess(ctx infrastructure.Context, obj response.DeleteContactResponse) error
	InfoSuccess(ctx infrastructure.Context, obj response.InfoContactResponse) error
}
