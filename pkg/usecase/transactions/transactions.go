package history

import (
	"time"

	domainAccount "anymind/pkg/domain/account"
	domainTransactions "anymind/pkg/domain/transactions"
	ctxSess "anymind/pkg/shared/utils/context"
	historySvc "anymind/pkg/usecase/history"
	"github.com/google/uuid"
)

type service struct {
	transactionsRepo domainTransactions.Repository
	accountRepo      domainAccount.Repository
	historyService   historySvc.Service
}

func NewService(transactionsRepo domainTransactions.Repository, accountRepo domainAccount.Repository, historyService historySvc.Service) Service {
	s := &service{
		transactionsRepo: transactionsRepo,
		accountRepo:      accountRepo,
		historyService:   historyService,
	}
	if s.transactionsRepo == nil {
		panic("please provide transaction repo")
	}
	if s.accountRepo == nil {
		panic("please provide account repo")
	}
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
	msg := &historySvc.HistoryTransaction{
		AccountUUID: account.ID,
		Amount:      req.Amount,
		DateTime:    req.DateTime.UTC(),
	}
	err = s.historyService.SaveHistory(ctxSess, msg)
	return
}
