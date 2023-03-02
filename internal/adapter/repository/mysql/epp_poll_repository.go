package mysql

import (
	"gitlab.com/merekmu/go-epp-rest/internal/adapter"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/entities"
	repository "gitlab.com/merekmu/go-epp-rest/internal/domain/repository/mysql"
)

type eppPollRepository struct {
	eppClient adapter.EppClient
}

func NewEppPollRepository(eppClient adapter.EppClient) repository.EppPollRepository {
	return &eppPollRepository{eppClient}
}

func (r *eppPollRepository) Insert(eppPoll entities.EPPPoll) error {
	return nil
}
