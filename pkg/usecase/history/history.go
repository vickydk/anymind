package history

import (
	"time"

	domainHistory "anymind/pkg/domain/history"
	ctxSess "anymind/pkg/shared/utils/context"
)

type service struct {
	historyRepo domainHistory.Repository
}

func NewService(historyRepo domainHistory.Repository) Service {
	s := &service{
		historyRepo: historyRepo,
	}
	if s.historyRepo == nil {
		panic("please provide history repo")
	}
	return s
}

func (s *service) SaveHistory(ctxSess *ctxSess.Context, req *HistoryTransaction) (err error) {
	lastHistory, err := s.historyRepo.FindByLatestTransaction(req.AccountUUID)
	if err != nil {
		//TODO: flag DB balance with error
		ctxSess.ErrorMessage = err.Error()
		return err
	}
	entity := &domainHistory.Entity{
		AccountUuid:  req.AccountUUID,
		DateTime:     req.DateTime.UTC(),
		Amount:       req.Amount,
		AmountBefore: lastHistory.AmountAfter,
		AmountAfter:  lastHistory.AmountAfter + req.Amount,
		CreatedAt:    time.Now().UTC(),
	}
	return s.historyRepo.Save(entity)
}
