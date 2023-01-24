package http

import (
	"net/http"

	"anymind/pkg/shared/constants"
	"anymind/pkg/shared/models"
	"github.com/labstack/echo/v4"
)

func errHandler(err error) (code int, message string) {
	if errResp, ok := err.(*constants.ApplicationError); ok {
		code = errResp.Code
		message = errResp.Message
	} else {
		return errHandler(constants.ErrorGeneral)
	}
	return
}

func httpError(c echo.Context, code int, err error) error {
	if err = c.JSON(code, map[string]string{"message": err.Error()}); err != nil {
		return err
	}
	return err
}

func httpOk(c echo.Context, out interface{}) error {
	resp := models.DefaultResponse{
		Data:    out,
		Message: "Success",
	}
	return c.JSON(http.StatusOK, resp)
}
