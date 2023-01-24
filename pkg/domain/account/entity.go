package transactions

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Entity) TableName() string {
	return "account"
}
