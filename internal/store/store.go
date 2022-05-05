package store

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Store struct {
		db *mongo.Database
	}
)

//NewStoriRepository return client db
func NewStoriRepository(ctx context.Context, uri string) (*Store, error) {
	mongo, err := newMongoConnection(ctx, uri)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: mongo,
	}, nil
}

// newMongoConnection connect to mongo DB
func newMongoConnection(ctx context.Context, uri string) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return &mongo.Database{}, entity.CustomError{
			Err:      fmt.Errorf("cannot connect to Database: %w", err),
			HTTPCode: http.StatusInternalServerError,
			Code:     "2b9d30c9-d5ee-4145-b925-0e11a7f1c2fb",
		}
	}

	mongoDatabase := os.Getenv("MONGO_DATABASE")

	db := client.Database(mongoDatabase)

	return db, nil
}
