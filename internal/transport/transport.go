package transport

import (
	"fmt"
	"net/http"
	"os"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Router struct {
		*echo.Echo
		Address string
		Handler EchoHandler
	}
	EchoHandler struct {
		StoriUseCases UseCases
	}
	UseCases struct {
		Stori StoriUsecase
	}
)

//StoriUsecase implement the methods of usecase (business logic)
type StoriUsecase interface {
	ValidateTransaction(transactions []entity.Transaction) error
}

//LocalHost routing
func (r *Router) LocalHost() error {
	base := r.Group("/custom-endpoints")
	user := os.Getenv("BASIC_AUTH_USER")
	pass := os.Getenv("BASIC_AUTH_PASSWORD")

	base.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == user && password == pass {
			return true, nil
		}
		return false, entity.CustomError{
			Err:      fmt.Errorf("basic auth failed"),
			HTTPCode: http.StatusUnauthorized,
			Code:     "e6807c42-3568-41de-a15f-fe0f073ab657",
		}
	}))
	transaction := base.Group("/transactions")
	transaction.GET("", r.Handler.Transactions)

	return r.Echo.Start(r.Address)
}
