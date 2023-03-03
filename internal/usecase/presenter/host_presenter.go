package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
)

type HostPresenter interface {
	Check(response []byte) (response.CheckHostResponse, error)
	Create(response []byte) (response.CreateHostResponse, error)
	Update(response []byte) (response.UpdateHostResponse, error)
	Delete(response []byte) (response.DeleteHostResponse, error)
	Info(response []byte) (response.InfoHostResponse, error)
}
