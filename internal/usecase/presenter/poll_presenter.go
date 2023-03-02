package presenter

import "gitlab.com/merekmu/go-epp-rest/internal/adapter/dto/response"

type PollPresenter interface {
	Acknowledge(responseObject response.PollAckResponse) string
	Request(responseObject response.PollRequestResponse) string
}
