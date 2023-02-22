package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
)

type HostPresenter interface {
	MapCheckResponse(response []byte) (model.CheckHostResponse, error)
	MapCreateResponse(response []byte) (model.CreateHostResponse, error)
}
