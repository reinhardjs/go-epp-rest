package mapper

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/entities"
)

type DtoToEntity interface {
	MapPollRequestResponseToEppPollEntity(input response.PollRequestResponse) (output entities.EPPPoll, err error)
}
