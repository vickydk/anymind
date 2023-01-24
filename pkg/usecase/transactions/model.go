package history

import (
	"time"
)

type TransactionRequest struct {
	AccountUUID string    `json:"account_uuid"`
	Amount      float64   `json:"amount"`
	DateTime    time.Time `json:"date_time"`
}
