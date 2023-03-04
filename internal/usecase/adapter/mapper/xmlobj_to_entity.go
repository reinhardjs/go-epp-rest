package mapper

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/entities"
)

type DtoToEntity interface {
	MapPollRequestResponseToEppPollEntity(input any) (output entities.EPPPoll)
}
