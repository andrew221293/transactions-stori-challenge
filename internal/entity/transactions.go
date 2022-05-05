package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	DataFraud struct {
		ID          string             `json:"id" bson:"_id"`
		Date        primitive.DateTime `bson:"date" json:"date"`
		Transaction string             `bson:"transaction" json:"transaction"`
		CreatedAt   primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
		UpdatedAt   primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
		DeletedAt   primitive.DateTime `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	}
)
