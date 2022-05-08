package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	User struct {
		ID        string             `json:"id" bson:"_id"`
		Name      string             `bson:"name" json:"name"`
		LastName  string             `bson:"last_name" json:"last_name"`
		Balance   float64            `bson:"balance" json:"balance"`
		Email     string             `bson:"email" json:"email"`
		CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
		UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
		DeletedAt primitive.DateTime `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	}
)
