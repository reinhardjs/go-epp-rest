package presenter

import (
	"fmt"

	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type pollPresenter struct{}

func NewPollPresenter() presenter.PollPresenter {
	return &pollPresenter{}
}

func (p *pollPresenter) Acknowledge(responseObject presenter.PollResponse) string {
	return fmt.Sprintf("%v %v", responseObject.GetResultCode(), responseObject.GetResultMessage())
}

func (p *pollPresenter) Request(responseObject presenter.PollResponse) string {
	return fmt.Sprintf("%v %v", responseObject.GetResultCode(), responseObject.GetResultMessage())
}
