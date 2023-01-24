package transactions

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID          uuid.UUID
	AccountUuid uuid.UUID
	Amount      float64
	FlagError   bool
	ErrorDetail string
	DateTime    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Entity) TableName() string {
	return "transactions"
}
