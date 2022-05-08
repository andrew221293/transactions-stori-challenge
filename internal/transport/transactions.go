package transport

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"

	"github.com/labstack/echo/v4"
)

// Transactions process and validate the CSV file
func (e EchoHandler) Transactions(c echo.Context) error {
	ctx := c.Request().Context()
	transactions, err := readCSVFile()
	if err != nil {
		return err
	}

	// send the transactions to the useCase
	transaction, err := e.StoriUseCases.Stori.ValidateTransaction(ctx, transactions)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transaction)

	return nil
}

// readCSVFile handler reading the CSV file
func readCSVFile() ([]entity.Transaction, error) {
	// open csv file
	csvFile, err := os.Open("txns.csv")
	if err != nil {
		return nil, entity.CustomError{
			Err:      err,
			HTTPCode: http.StatusBadRequest,
			Code:     "17b30b8e-99a6-424e-82ab-b39d21eb532af",
		}
	}

	// read csv file
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, entity.CustomError{
			Err:      err,
			HTTPCode: http.StatusBadRequest,
			Code:     "8797a4d7-1d4d-42a5-9ce0-99d3992115be",
		}
	}

	var transactionData []entity.Transaction
	// assign the lines of CSV files to my struct
	for i, line := range csvLines {
		if i != 0 {
			idTransaction, err := strconv.Atoi(line[0])
			if err != nil {
				return nil, entity.CustomError{
					Err:      err,
					HTTPCode: http.StatusBadRequest,
					Code:     "f4220fb8-a059-4dbe-9dfe-a3597f47cb0e",
				}
			}
			emp := entity.Transaction{
				ID:          idTransaction,
				Date:        line[1],
				Transaction: line[2],
			}
			transactionData = append(transactionData, emp)
		}
	}

	defer csvFile.Close()

	return transactionData, nil
}
