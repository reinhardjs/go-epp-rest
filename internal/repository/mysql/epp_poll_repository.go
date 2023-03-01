package mysql

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/entities"
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter"
	repository "gitlab.com/merekmu/go-epp-rest/internal/usecase/repository/mysql"
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
