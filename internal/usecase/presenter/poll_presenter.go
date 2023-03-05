package presenter

import "gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"

type PollPresenter interface {
	Poll(responseObject response.PollRequestResponse) string
}
