package history

import "github.com/google/uuid"

type Repository interface {
	Save(entity *Entity) (err error)
	FindByLatestTransaction(AccountUUID uuid.UUID) (entity Entity, err error)
}
