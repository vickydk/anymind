package gorm

import (
	"errors"
	"time"

	domainHistory "anymind/pkg/domain/history"
	"anymind/pkg/shared/constants"
	"anymind/pkg/shared/database"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type historyRepository struct {
	dbMaster *database.Database
	dbSlave  *database.Database
}

func HistorySetup(dbMaster *database.Database, dbSlave *database.Database) *historyRepository {
	r := &historyRepository{dbMaster: dbMaster, dbSlave: dbSlave}
	if r.dbMaster == nil {
		panic("please provide db master")
	}
	if r.dbSlave == nil {
		panic("please provide db slave")
	}
	return r
}

func (r *historyRepository) Save(entity *domainHistory.Entity) (err error) {
	entity.UpdatedAt = time.Now().UTC()
	err = r.dbMaster.Save(entity).Error
	if err != nil {
		return
	}

	return
}

func (r *historyRepository) FindByLatestTransaction(AccountUUID uuid.UUID) (entity domainHistory.Entity, err error) {
	err = r.dbSlave.
		Where("account_uuid = ?", AccountUUID).
		Order("id desc").
		First(&entity).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity, nil
		}
		err = constants.ErrorDatabase
		return
	}

	return
}
