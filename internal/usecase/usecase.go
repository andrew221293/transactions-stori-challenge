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
	InsertUser(ctx context.Context, user entity.User) error
}
