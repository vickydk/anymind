package history

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID           int64
	AccountUuid  uuid.UUID
	Amount       float64
	AmountBefore float64
	AmountAfter  float64
	DateTime     time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (t *Entity) TableName() string {
	return "history"
}
