package queue

import "anymind/pkg/interface/container"

type Handler struct {
	history *handlerHistory
}

func SetupHandlers(container *container.Container) *Handler {
	history := newHandlerHistory(container.HistorySvc)

	return &Handler{
		history: history,
	}
}
