package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
)

type ContactPresenter interface {
	MapCheckResponse(response []byte) (model.CheckContactResponse, error)
	MapCreateResponse(response []byte) (model.CreateContactResponse, error)
}
