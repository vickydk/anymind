package history

import ctxSess "anymind/pkg/shared/utils/context"

type Service interface {
	SaveHistory(ctxSess *ctxSess.Context, req *HistoryTransaction) (err error)
	SearchHistory(ctxSess *ctxSess.Context, req *SearchHistory) (resp []*History, err error)
}
