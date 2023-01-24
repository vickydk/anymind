package history

import (
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Save(entity *Entity) (err error)
	FindByLatestTransaction(AccountUUID uuid.UUID) (entity Entity, err error)
	FindByDateTime(startDateTime, endDateTime time.Time) (entity []*Entity, err error)
}
