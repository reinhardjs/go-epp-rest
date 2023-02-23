package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
)

type ContactPresenter interface {
	MapCheckResponse(response []byte) (model.CheckContactResponse, error)
	MapCreateResponse(response []byte) (model.CreateContactResponse, error)
	MapUpdateResponse(response []byte) (model.UpdateContactResponse, error)
	MapDeleteResponse(response []byte) (model.DeleteContactResponse, error)
	MapInfoResponse(response []byte) (model.InfoContactResponse, error)
}
