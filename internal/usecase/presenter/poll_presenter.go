package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/common/dto/response"
)

type PollPresenter interface {
	Acknowledge(responseObject PollResponse) string
	Request(responseObject PollResponse) string
}

type PollResponse interface {
	GetResultCode() int
	GetResultMessage() string
}

type PollRequestResponseImpl struct {
	DTO *response.PollRequestResponse
}

func (response *PollRequestResponseImpl) GetResultCode() int {
	return response.DTO.Result.Code
}

func (response *PollRequestResponseImpl) GetResultMessage() string {
	return response.DTO.Result.Message
}
