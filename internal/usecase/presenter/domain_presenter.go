package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter/dto/response"
)

type DomainPresenter interface {
	Check(response []byte) (response.CheckDomainResponse, error)
	Create(response []byte) (response.CreateDomainResponse, error)
	Delete(response []byte) (response.DeleteDomainResponse, error)
	Info(response []byte) (response.InfoDomainResponse, error)
	SecDNSUpdate(response []byte) (response.SecDNSUpdateResponse, error)
}
