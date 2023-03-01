package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter/dto/response"
)

type PollPresenter interface {
	Acknowledge(response []byte) (response.PollAckResponse, error)
	Request(response []byte) (response.PollRequestResponse, error)
}
