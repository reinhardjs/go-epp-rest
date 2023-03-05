package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
)

type HostPresenter interface {
	Check(response.CheckHostResponse) string
	Create(response.CreateHostResponse) string
	Update(response.UpdateHostResponse) string
	Delete(response.DeleteHostResponse) string
	Info(response.InfoHostResponse) string
}
