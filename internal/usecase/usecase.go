package usecase

import (
	"context"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"
)

type (
	StoriUseCase struct {
		Store StoriStore
	}
)

//StoriStore Implement all database methods
type StoriStore interface {
	InserTransactionHistory(ctx context.Context, transaction entity.TransactionHistory) error
}
