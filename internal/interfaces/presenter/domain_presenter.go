package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
)

type DomainPresenter interface {
	Check(response response.CheckDomainResponse) string
	Create(response response.CreateDomainResponse) string
	Delete(response response.DeleteDomainResponse) string
	Info(response response.InfoDomainResponse) string
	SecDNSUpdate(response response.SecDNSUpdateResponse) string
}
