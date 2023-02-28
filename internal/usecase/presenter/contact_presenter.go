package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/responses"
)

type ContactPresenter interface {
	MapCheckResponse(response []byte) (responses.CheckContactResponse, error)
	MapCreateResponse(response []byte) (responses.CreateContactResponse, error)
	MapUpdateResponse(response []byte) (responses.UpdateContactResponse, error)
	MapDeleteResponse(response []byte) (responses.DeleteContactResponse, error)
	MapInfoResponse(response []byte) (responses.InfoContactResponse, error)
}
