package mysql

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/entities"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure"
	repository "gitlab.com/merekmu/go-epp-rest/internal/usecase/repository/mysql"
)

type eppPollRepository struct {
	eppClient infrastructure.EppClient
}

func NewEppPollRepository(eppClient infrastructure.EppClient) repository.EppPollRepository {
	return &eppPollRepository{eppClient}
}

func (r *eppPollRepository) Insert(eppPoll entities.EPPPoll) error {
	return nil
}
