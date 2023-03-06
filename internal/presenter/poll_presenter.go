package presenter

import (
	"fmt"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type pollPresenter struct{}

func NewPollPresenter() presenter.PollPresenter {
	return &pollPresenter{}
}

func (p *pollPresenter) PollSuccess(ctx infrastructure.Context, obj response.PollRequestResponse) {
	var message string

	if obj.Result.Code == 1301 {
		message = fmt.Sprintf("%v %v", 1000, "Command Completed Successfully")
	}

	if obj.Result.Code == 1300 {
		message = fmt.Sprintf("%v %v", 1000, "No Message")
	}

	ctx.String(200, message)
}
