package queue

import ctxSess "anymind/pkg/shared/utils/context"

type Wrapper interface {
	SendHistoryTransaction(ctxSess *ctxSess.Context, in *HistoryTransaction) (err error)
}
