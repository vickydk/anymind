package history

import ctxSess "anymind/pkg/shared/utils/context"

type Service interface {
	AddTransaction(ctxSess *ctxSess.Context, req *TransactionRequest) (err error)
}
