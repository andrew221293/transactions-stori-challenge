package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"

	"github.com/labstack/echo/v4"
)

func (e EchoHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	//parsing body
	var user entity.User
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return entity.CustomError{
			Err:      fmt.Errorf("error parsing body"),
			HTTPCode: http.StatusBadRequest,
			Code:     "5e4ec734-bf45-46e3-942c-6ae24fea3d04",
		}
	}

	err = validateBody(user)
	if err != nil {
		return err
	}

	response, err := e.StoriUseCases.Stori.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func validateBody(user entity.User) error {
	if user.Name == "" {
		return entity.CustomError{
			Err:      fmt.Errorf("a name is needed"),
			HTTPCode: http.StatusBadRequest,
			Code:     "4a4113b2-1d9d-4be3-94a8-c6b9255285ee",
		}
	}

	if user.LastName == "" {
		return entity.CustomError{
			Err:      fmt.Errorf("a lastname is needed"),
			HTTPCode: http.StatusBadRequest,
			Code:     "9844d655-ba6f-4f07-951c-604a653527b3",
		}
	}

	if user.Balance == 0 {
		return entity.CustomError{
			Err:      fmt.Errorf("a balance is needed"),
			HTTPCode: http.StatusBadRequest,
			Code:     "a8e6a88d-04ce-46ab-b758-f8ca812dc70f",
		}
	}

	return nil
}
