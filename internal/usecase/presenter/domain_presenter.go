package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/responses"
)

type DomainPresenter interface {
	MapCheckResponse(response []byte) (responses.CheckDomainResponse, error)
	MapCreateResponse(response []byte) (responses.CreateDomainResponse, error)
	MapDeleteResponse(response []byte) (responses.DeleteDomainResponse, error)
	MapInfoResponse(response []byte) (responses.InfoDomainResponse, error)
	MapSecDNSUpdateResponse(response []byte) (responses.SecDNSUpdateResponse, error)
}
