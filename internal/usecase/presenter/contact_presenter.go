package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
)

type ContactPresenter interface {
	MapResponse(response []byte) (model.CheckContactResponse, error)
}
