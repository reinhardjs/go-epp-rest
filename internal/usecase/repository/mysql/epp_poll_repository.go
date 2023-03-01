package mysql

import entities "gitlab.com/merekmu/go-epp-rest/internal/model/domain"

type EppPollRepository interface {
	Insert(eppPoll entities.EPPPoll) error
}
