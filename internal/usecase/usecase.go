package usecase

import "gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"

type TransferInteractor interface {
	Check(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Request(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Cancel(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Approve(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Reject(ctx infrastructure.Context, data interface{}, ext string, langTag string)
}

type HostInteractor interface {
	Check(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Create(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Update(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Info(ctx infrastructure.Context, data interface{}, ext string, langTag string)
}

type DomainInteractor interface {
	Check(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Create(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	Info(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	SecDNSUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string)
	ContactUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string)
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
