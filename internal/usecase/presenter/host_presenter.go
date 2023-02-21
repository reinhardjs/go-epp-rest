package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
)

type HostPresenter interface {
	MapResponse(response []byte) (model.CheckHostResponse, error)
}
