package presenter

import (
	"fmt"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type pollPresenter struct{}

func NewPollPresenter() presenter.PollPresenter {
	return &pollPresenter{}
}

func (p *pollPresenter) Poll(responseObject response.PollRequestResponse) (message string) {

	if responseObject.Result.Code == 1301 {
		message = fmt.Sprintf("%v %v", 1000, "Command Completed Successfully")
	}

	if responseObject.Result.Code == 1300 {
		message = fmt.Sprintf("%v %v", 1000, "No Message")
	}

	return
}
