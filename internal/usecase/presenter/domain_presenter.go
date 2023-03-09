package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
)

type DomainPresenter interface {
	CheckSuccess(ctx infrastructure.Context, obj response.CheckDomainResponse) error
	CreateSuccess(ctx infrastructure.Context, obj response.CreateDomainResponse) error
	DeleteSuccess(ctx infrastructure.Context, obj response.DeleteDomainResponse) error
	InfoSuccess(ctx infrastructure.Context, obj response.InfoDomainResponse) error
	SecDNSUpdateSuccess(ctx infrastructure.Context, obj response.SecDNSUpdateResponse) error
	ContactUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse) error
	StatusUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse) error
	AuthInfoUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse) error
	NameserverUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse) error
	RenewSuccess(ctx infrastructure.Context, obj response.RenewDomainResponse) error
}
