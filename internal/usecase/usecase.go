package usecase

import (
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type TransferInteractor interface {
	Check(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Request(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Cancel(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Approve(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Reject(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
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
	Check(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Create(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Info(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	SecDNSUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	ContactUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	StatusUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	AuthInfoUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	NameserverUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Renew(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
}

type ContactInteractor interface {
	Check(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Create(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Update(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
	Info(ctx infrastructure.Context, data interface{}, ext string, langTag string) error
}

type PollInteractor interface {
	Poll(ctx infrastructure.Context) error
}
