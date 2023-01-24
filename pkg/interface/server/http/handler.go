package http

import "anymind/pkg/interface/container"

type Handler struct {
	transactionHandler *transactionsHandler
}

func SetupHandlers(container *container.Container) *Handler {
	transactionHandler := SetupTransactionsHandler(container.Validate, container.TransactionSvc)
	return &Handler{
		transactionHandler: transactionHandler,
	}
}
