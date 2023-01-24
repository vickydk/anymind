package queue

import (
	"anymind/pkg/shared/config"
	Queue "anymind/pkg/shared/queue"
	ctxSess "anymind/pkg/shared/utils/context"
)

type kafkaQueue struct {
	q   Queue.Producer
	cfg *config.KafkaProcedureTopics
}

func SetupQueue(q Queue.Producer, cfg *config.KafkaProcedureTopics) *kafkaQueue {
	if q == nil {
		panic("queue client is nil")
	}
	if cfg == nil {
		panic("trx pending config is nil")
	}
	return &kafkaQueue{
		q:   q,
		cfg: cfg,
	}
}

func (p *kafkaQueue) SendHistoryTransaction(ctxSess *ctxSess.Context, in *HistoryTransaction) (err error) {
	return p.q.SendMessage(ctxSess.XRequestID, p.cfg.HistoryTransaction, in)
}
