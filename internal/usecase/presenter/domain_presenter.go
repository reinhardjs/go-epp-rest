package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
)

type DomainPresenter interface {
	MapCheckResponse(response []byte) (model.CheckDomainResponse, error)
	MapCreateResponse(response []byte) (model.CreateDomainResponse, error)
	MapDeleteResponse(response []byte) (model.DeleteDomainResponse, error)
	MapInfoResponse(response []byte) (model.InfoDomainResponse, error)
	MapSecDNSUpdateResponse(response []byte) (model.SecDNSUpdateResponse, error)
}
