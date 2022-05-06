package store

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"
)

func (s StoriStore) InsertUser(ctx context.Context, user entity.User) error {
	collection := s.db.Collection("users")

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return entity.CustomError{
			Err:      fmt.Errorf("cannot save the data: %w", err),
			HTTPCode: http.StatusInternalServerError,
			Code:     "df48e536-eb81-4d16-903a-892cf6d81175",
		}
	}

	return nil
}
