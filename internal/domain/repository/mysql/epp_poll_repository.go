package mysql

import "gitlab.com/merekmu/go-epp-rest/internal/domain/entities"

type EppPollRepository interface {
	Insert(eppPoll entities.EPPPoll) error
}
