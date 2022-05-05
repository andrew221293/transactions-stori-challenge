package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (e EchoHandler) Transactions(c echo.Context) error {

	return c.JSON(http.StatusOK, "Hi Transport")

	return nil
}
