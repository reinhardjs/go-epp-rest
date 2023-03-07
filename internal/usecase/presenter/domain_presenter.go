package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
)

type DomainPresenter interface {
	CheckSuccess(ctx infrastructure.Context, obj response.CheckDomainResponse)
	CreateSuccess(ctx infrastructure.Context, obj response.CreateDomainResponse)
	DeleteSuccess(ctx infrastructure.Context, obj response.DeleteDomainResponse)
	InfoSuccess(ctx infrastructure.Context, obj response.InfoDomainResponse)
	SecDNSUpdateSuccess(ctx infrastructure.Context, obj response.SecDNSUpdateResponse)
	ContactUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse)
	StatusUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse)
	AuthInfoUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse)
}
