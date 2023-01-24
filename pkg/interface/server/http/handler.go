package http

import "anymind/pkg/interface/container"

type Handler struct {
	transactionHandler *transactionsHandler
	historyHandler     *historyHandler
}

func SetupHandlers(container *container.Container) *Handler {
	transactionHandler := SetupTransactionsHandler(container.Validate, container.TransactionSvc)
	historyHandler := SetupHistoryHandler(container.Validate, container.HistorySvc)
	return &Handler{
		transactionHandler: transactionHandler,
		historyHandler:     historyHandler,
	}
}
