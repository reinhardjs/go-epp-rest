package presenter

import "gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter/dto/response"

type PollPresenter interface {
	Acknowledge(responseObject response.PollAckResponse) string
	Request(responseObject response.PollRequestResponse) string
}
