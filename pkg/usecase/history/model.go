package history

import (
	"time"

	"github.com/google/uuid"
)

type HistoryTransaction struct {
	AccountUUID uuid.UUID `json:"account_uuid" validate:"required"`
	Amount      float64   `json:"amount" validate:"required"`
	DateTime    time.Time `json:"date_time" validate:"required"`
}

type SearchHistory struct {
	StartDateTime time.Time `json:"startDateTime" validate:"required"`
	EndDateTime   time.Time `json:"endDateTime" validate:"required"`
}

type History struct {
	Amount   float64   `json:"amount"`
	DateTime time.Time `json:"dateTime"`
}
