package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
)

type ContactPresenter interface {
	Check(response.CheckContactResponse) string
	Create(response.CreateContactResponse) string
	Update(response.UpdateContactResponse) string
	Delete(response.DeleteContactResponse) string
	Info(response.InfoContactResponse) string
}
