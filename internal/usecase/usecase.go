package usecase

type TransferInteractor interface {
	Check(data interface{}, ext string, langTag string) (res string, err error)
	Request(data interface{}, ext string, langTag string) (res string, err error)
	Cancel(data interface{}, ext string, langTag string) (res string, err error)
	Approve(data interface{}, ext string, langTag string) (res string, err error)
	Reject(data interface{}, ext string, langTag string) (res string, err error)
}

type HostInteractor interface {
	Check(data interface{}, ext string, langTag string) (res string, err error)
	Create(data interface{}, ext string, langTag string) (res string, err error)
	Update(data interface{}, ext string, langTag string) (res string, err error)
	Delete(data interface{}, ext string, langTag string) (res string, err error)
	Info(data interface{}, ext string, langTag string) (res string, err error)
}

type DomainInteractor interface {
	Check(data interface{}, ext string, langTag string) (res string, err error)
	Create(data interface{}, ext string, langTag string) (res string, err error)
	Delete(data interface{}, ext string, langTag string) (res string, err error)
	Info(data interface{}, ext string, langTag string) (res string, err error)
	SecDNSUpdate(data interface{}, ext string, langTag string) (res string, err error)
}

type ContactInteractor interface {
	Check(data interface{}, ext string, langTag string) (res string, err error)
	Create(data interface{}, ext string, langTag string) (res string, err error)
	Update(data interface{}, ext string, langTag string) (res string, err error)
	Delete(data interface{}, ext string, langTag string) (res string, err error)
	Info(data interface{}, ext string, langTag string) (res string, err error)
}

type PollInteractor interface {
	Poll() (res string, err error)
}
