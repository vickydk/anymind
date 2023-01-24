package gorm

import (
	"errors"
	"time"

	domainTransactions "anymind/pkg/domain/transactions"
	"anymind/pkg/shared/constants"
	"anymind/pkg/shared/database"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type transactionsRepository struct {
	dbMaster *database.Database
	dbSlave  *database.Database
}

func TransactionsSetup(dbMaster *database.Database, dbSlave *database.Database) *transactionsRepository {
	r := &transactionsRepository{dbMaster: dbMaster, dbSlave: dbSlave}
	if r.dbMaster == nil {
		panic("please provide db master")
	}
	if r.dbSlave == nil {
		panic("please provide db slave")
	}
	return r
}

func (r *transactionsRepository) Save(entity *domainTransactions.Entity) (err error) {
	entity.UpdatedAt = time.Now().UTC()
	err = r.dbMaster.Save(entity).Error
	if err != nil {
		return
	}

	return
}

func (r *transactionsRepository) FindById(id uuid.UUID) (entity domainTransactions.Entity, err error) {
	err = r.dbSlave.
		Where("id = ?", id).
		First(&entity).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = constants.ErrorAccountNotFound
			return
		}
		err = constants.ErrorDatabase
		return
	}

	return
}
