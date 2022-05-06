package store

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
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

func (s StoriStore) GetOneUser(ctx context.Context, userId string) (entity.User, error) {
	var user entity.User
	collection := s.db.Collection("users")

	if err := collection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user); err != nil {
		return user, entity.CustomError{
			Err:      fmt.Errorf("cannot get the data: %w", err),
			HTTPCode: http.StatusInternalServerError,
			Code:     "c0d94606-4366-49c3-b020-d96dab1bb6f1",
		}
	}

	return user, nil
}
