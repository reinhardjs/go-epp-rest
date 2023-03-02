package repository

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/entities"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/repository"
	"gorm.io/gorm"
)

type eppPollRepository struct {
	dbConn *gorm.DB
}

func NewEppPollRepository(dbConn *gorm.DB) repository.EppPollRepository {
	return &eppPollRepository{dbConn}
}

func (r *eppPollRepository) Insert(eppPoll entities.EPPPoll) error {
	return nil
}
