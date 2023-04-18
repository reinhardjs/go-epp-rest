package repository

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/entities"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
	"gorm.io/gorm"
)

type eppPollRepository struct {
	dbConn *gorm.DB
}

func NewEppPollRepository(dbConn *gorm.DB) repository.EppPollRepository {
	return &eppPollRepository{dbConn}
}

func (r *eppPollRepository) Insert(eppPoll entities.EPPPoll) error {
	tx := r.dbConn.Create(&eppPoll)

	if tx.Error != nil {
		return errors.Wrap(tx.Error, "EppPollRepository Insert: r.dbConn.Create")
	}

	return nil
}
