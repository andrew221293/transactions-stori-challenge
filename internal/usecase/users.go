package usecase

import (
	"context"
	"time"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s StoriUseCase) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	user.ID = primitive.NewObjectID().Hex()
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	err := s.Store.InsertUser(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
