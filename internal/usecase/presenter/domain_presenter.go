package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
)

type DomainPresenter interface {
	MapResponse(response []byte) (model.CheckDomainResponse, error)
}
