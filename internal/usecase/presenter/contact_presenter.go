package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter/dto/response"
)

type ContactPresenter interface {
	Check(response []byte) (response.CheckContactResponse, error)
	Create(response []byte) (response.CreateContactResponse, error)
	Update(response []byte) (response.UpdateContactResponse, error)
	Delete(response []byte) (response.DeleteContactResponse, error)
	Info(response []byte) (response.InfoContactResponse, error)
}
