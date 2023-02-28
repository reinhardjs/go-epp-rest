package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/responses"
)

type HostPresenter interface {
	MapCheckResponse(response []byte) (responses.CheckHostResponse, error)
	MapCreateResponse(response []byte) (responses.CreateHostResponse, error)
	MapUpdateResponse(response []byte) (responses.UpdateHostResponse, error)
	MapDeleteResponse(response []byte) (responses.DeleteHostResponse, error)
	MapInfoResponse(response []byte) (responses.InfoHostResponse, error)
}
