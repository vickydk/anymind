package gorm

import (
	"errors"
	"time"

	domainAccount "anymind/pkg/domain/account"
	"anymind/pkg/shared/constants"
	"anymind/pkg/shared/database"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type accountRepository struct {
	dbMaster *database.Database
	dbSlave  *database.Database
}

func AccountSetup(dbMaster *database.Database, dbSlave *database.Database) *accountRepository {
	r := &accountRepository{dbMaster: dbMaster, dbSlave: dbSlave}
	if r.dbMaster == nil {
		panic("please provide db master")
	}
	if r.dbSlave == nil {
		panic("please provide db slave")
	}
	return r
}

func (r *accountRepository) Save(entity *domainAccount.Entity) (err error) {
	entity.UpdatedAt = time.Now().UTC()
	err = r.dbMaster.Save(entity).Error
	if err != nil {
		return
	}

	return
}

func (r *accountRepository) FindById(id uuid.UUID) (entity domainAccount.Entity, err error) {
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
