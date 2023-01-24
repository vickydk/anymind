package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetupRouter(server *echo.Echo, handler *Handler) {
	server.GET("/ping", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "services up and running... "+time.Now().Format(time.RFC3339))
	})
	server.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "OK")
	})

	root := server.Group("/api/v1")
	root.POST("/transactions", handler.transactionHandler.addTransactions)

}
