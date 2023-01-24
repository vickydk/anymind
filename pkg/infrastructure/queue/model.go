package queue

import (
	"time"

	"github.com/google/uuid"
)

type HistoryTransaction struct {
	AccountUUID uuid.UUID `json:"account_uuid"`
	Amount      float64   `json:"amount"`
	DateTime    time.Time `json:"date_time"`
}
