package queue

import (
	"encoding/json"

	Logger "anymind/pkg/shared/logger"
	CtxSess "anymind/pkg/shared/utils/context"
	historySvc "anymind/pkg/usecase/history"
)

type handlerHistory struct {
	service historySvc.Service
}

func newHandlerHistory(service historySvc.Service) *handlerHistory {
	if service == nil {
		panic("service is nil")
	}
	return &handlerHistory{
		service: service,
	}
}

func (h *handlerHistory) HistoryTransaction(threadID, topic string, msg []byte) {
	sess := h.createSession(threadID, topic)
	sess.Lv1("incoming Request")

	request := &historySvc.HistoryTransaction{}
	if err := json.Unmarshal(msg, request); err != nil {
		sess.ErrorMessage = err.Error()
		sess.Lv4("failed to unmarshal request")
		return
	}

	sess.Request = request

	err := h.service.SaveHistory(sess, request)

	sess.Lv4(err)
}

func (h *handlerHistory) createSession(threadID, topic string) *CtxSess.Context {
	return CtxSess.New(Logger.GetLogger()).
		SetAppName("anymind.Queue").
		SetAppVersion("0.0").
		SetURL(topic).
		SetMethod("KAFKA").
		SetXRequestID(threadID)
}
