package history

import (
	"sync"
	"time"

	domainHistory "anymind/pkg/domain/history"
	ctxSess "anymind/pkg/shared/utils/context"
)

type service struct {
	historyRepo domainHistory.Repository
	mx          sync.Mutex
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
	s.mx.Lock()

	lastHistory, err := s.historyRepo.FindByLatestTransaction(req.AccountUUID)
	if err != nil {
		//TODO: flag DB balance with error
		ctxSess.ErrorMessage = err.Error()
		s.mx.Unlock()
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
	err = s.historyRepo.Save(entity)
	s.mx.Unlock()
	return
}

func (s *service) SearchHistory(ctxSess *ctxSess.Context, req *SearchHistory) (resp []*History, err error) {
	ls, err := s.historyRepo.FindByDateTime(req.StartDateTime, req.EndDateTime)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return
	}

	resp = []*History{}
	for _, v := range ls {
		resp = append(resp, &History{
			DateTime: v.DateTime,
			Amount:   v.AmountAfter,
		})
	}

	return
}
