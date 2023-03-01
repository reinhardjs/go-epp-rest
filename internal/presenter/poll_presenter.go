package presenter

import (
	"fmt"

	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type pollPresenter struct{}

func NewPollPresenter() presenter.PollPresenter {
	return &pollPresenter{}
}

func (p *pollPresenter) Acknowledge(responseObject response.PollAckResponse) string {
	return fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)
}

func (p *pollPresenter) Request(responseObject response.PollRequestResponse) string {
	return fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)
}
