package history

import ctxSess "anymind/pkg/shared/utils/context"

type Service interface {
	SaveHistory(ctxSess *ctxSess.Context, req *HistoryTransaction) (err error)
}
