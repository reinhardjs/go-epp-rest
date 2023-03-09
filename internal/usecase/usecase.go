package usecase

import (
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type TransferInteractor interface {
	Check(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Request(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Cancel(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Approve(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Reject(ctx infrastructure.Context, data interface{}, ext string, langTag string)
}

type HostInteractor interface {
	Check(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Create(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Update(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Info(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Change(ctx infrastructure.Context, data types.HostUpdateType, ext string, langTag string) error
	CheckAndCreate(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
}

type DomainInteractor interface {
	Check(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Create(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Info(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	SecDNSUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	ContactUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	StatusUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	AuthInfoUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	NameserverUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Renew(ctx infrastructure.Context, data interface{}, ext string, langTag string)
}

type ContactInteractor interface {
	Check(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Create(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Update(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Info(ctx infrastructure.Context, data interface{}, ext string, langTag string)
}

type PollInteractor interface {
	Poll(ctx infrastructure.Context)
}
