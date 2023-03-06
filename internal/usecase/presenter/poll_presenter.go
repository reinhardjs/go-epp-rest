package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
)

type PollPresenter interface {
	PollSuccess(ctx infrastructure.Context, obj response.PollRequestResponse)
}
