package http

import (
	"errors"
	"fmt"
	"net/http"

	historySvc "anymind/pkg/usecase/history"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type historyHandler struct {
	validate *validator.Validate
	service  historySvc.Service
}

func SetupHistoryHandler(validate *validator.Validate, service historySvc.Service) *historyHandler {
	handler := &historyHandler{
		validate: validate,
		service:  service,
	}
	if handler.service == nil {
		panic("service is nil")
	}
	return handler
}

func (s *historyHandler) searchHistory(c echo.Context) error {
	context := Parse(c)
	ctxSess := context.CtxSess
	request := &historySvc.SearchHistory{}
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
	resp, err := s.service.SearchHistory(ctxSess, request)
	if err != nil {
		ctxSess.Lv4()
		httpCode, errMsg := errHandler(err)
		return httpError(c, httpCode, errors.New(errMsg))
	}

	ctxSess.Lv4(resp)
	return httpOk(c, resp)
}
