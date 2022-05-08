package store

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"
)

func (s StoriStore) InserTransactionHistory(ctx context.Context, transaction entity.TransactionHistory) error {
	collection := s.db.Collection("transactions")

	_, err := collection.InsertOne(ctx, transaction)
	if err != nil {
		return entity.CustomError{
			Err:      fmt.Errorf("cannot save the data: %w", err),
			HTTPCode: http.StatusInternalServerError,
			Code:     "f658e4cd-ae87-4dde-baf3-9de5a06ffd9c",
		}
	}

	return nil
}
