package history

import (
	"sync"
	"time"

	domainAccount "anymind/pkg/domain/account"
	domainTransactions "anymind/pkg/domain/transactions"
	QueueWrapper "anymind/pkg/infrastructure/queue"
	Logger "anymind/pkg/shared/logger"
	"anymind/pkg/shared/utils"
	ctxSess "anymind/pkg/shared/utils/context"
	historySvc "anymind/pkg/usecase/history"
	"github.com/google/uuid"
)

type service struct {
	queue            QueueWrapper.Wrapper
	transactionsRepo domainTransactions.Repository
	accountRepo      domainAccount.Repository
	historyService   historySvc.Service

	msgQueue chan historySvc.HistoryTransaction
	quit     chan bool
	mx       sync.Mutex
}

func NewService(q QueueWrapper.Wrapper, transactionsRepo domainTransactions.Repository, accountRepo domainAccount.Repository, historyService historySvc.Service) Service {
	s := &service{
		queue:            q,
		transactionsRepo: transactionsRepo,
		accountRepo:      accountRepo,
		historyService:   historyService,
	}
	if s.queue == nil {
		panic("please provide queue wrapper")
	}
	if s.transactionsRepo == nil {
		panic("please provide transaction repo")
	}
	if s.accountRepo == nil {
		panic("please provide account repo")
	}
	s.msgQueue = make(chan historySvc.HistoryTransaction, 1)
	go s.scanCampaigns()
	return s
}

func (s *service) AddTransaction(ctxSess *ctxSess.Context, req *TransactionRequest) (err error) {
	accountID, err := uuid.Parse(req.AccountUUID)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return err
	}
	account, err := s.accountRepo.FindById(accountID)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return err
	}
	dateTime := time.Now().UTC()
	entity := &domainTransactions.Entity{
		AccountUuid: account.ID,
		Amount:      req.Amount,
		DateTime:    req.DateTime.UTC(),
		CreatedAt:   dateTime,
		UpdatedAt:   dateTime,
	}
	err = s.transactionsRepo.Save(entity)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return err
	}
	//msg := &QueueWrapper.HistoryTransaction{
	//	AccountUUID: account.ID,
	//	Amount:      req.Amount,
	//	DateTime:    req.DateTime.UTC(),
	//}
	//err = s.queue.SendHistoryTransaction(ctxSess, msg)
	msg := historySvc.HistoryTransaction{
		AccountUUID: account.ID,
		Amount:      req.Amount,
		DateTime:    req.DateTime.UTC(),
	}
	s.msgQueue <- msg
	return
}

func (s *service) scanCampaigns() {
	for {
		select {
		// Periodically scan the data source for campaigns to process.
		case e, ok := <-s.msgQueue:
			if !ok {
				return
			}
			s.mx.Lock()
			ctxSess := s.createSession()
			ctxSess.Lv1("incoming queue")
			if err := s.historyService.SaveHistory(ctxSess, &e); err != nil {
				ctxSess.ErrorMessage = err.Error()
			}
			ctxSess.Lv4()
			s.mx.Unlock()
		case <-s.quit:
			return
		}
	}
}

func (s *service) createSession() *ctxSess.Context {
	return ctxSess.New(Logger.GetLogger()).
		SetAppName("anymind.Queue").
		SetAppVersion("0.0").
		SetMethod("Queue").
		SetXRequestID(utils.GenerateThreadId())
}
