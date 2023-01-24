package http

import (
	"errors"
	"fmt"
	"net/http"

	transactionSvc "anymind/pkg/usecase/transactions"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type transactionsHandler struct {
	validate *validator.Validate
	service  transactionSvc.Service
}

func SetupTransactionsHandler(validate *validator.Validate, service transactionSvc.Service) *transactionsHandler {
	handler := &transactionsHandler{
		validate: validate,
		service:  service,
	}
	if handler.service == nil {
		panic("service is nil")
	}
	return handler
}

func (s *transactionsHandler) addTransactions(c echo.Context) error {
	context := Parse(c)
	ctxSess := context.CtxSess
	request := &transactionSvc.TransactionRequest{}
	if err := c.Bind(request); err != nil {
		ctxSess.ErrorMessage = err.Error()
		ctxSess.Lv4()
		return httpError(c, http.StatusBadRequest, fmt.Errorf("bind request: %w", err))
	}
	if err := s.validate.Struct(request); err != nil {
		ctxSess.ErrorMessage = err.Error()
		ctxSess.Lv4()
		return httpError(c, http.StatusBadRequest, fmt.Errorf("validate: %w", err))
	}

	ctxSess.Request = request
	err := s.service.AddTransaction(ctxSess, request)
	if err != nil {
		ctxSess.Lv4()
		httpCode, errMsg := errHandler(err)
		return httpError(c, httpCode, errors.New(errMsg))
	}

	ctxSess.Lv4()
	return httpOk(c, nil)
}
