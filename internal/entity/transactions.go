package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Transaction struct {
		ID          int    `json:"id" bson:"id"`
		UserID      string `json:"user_id" bson:"_user_id"`
		Date        string `bson:"date" json:"date"`
		Transaction string `bson:"transaction" json:"transaction"`
	}
	TransactionsPerMonth struct {
		month string
		total int
	}
	User struct {
		ID        string             `json:"id" bson:"_id"`
		Name      string             `bson:"name" json:"name"`
		LastName  string             `bson:"last_name" json:"last_name"`
		Balance   float64            `bson:"balance" json:"balance"`
		CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
		UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
		DeletedAt primitive.DateTime `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	}
	Balance struct {
		ID           string                 `json:"id" bson:"_id"`
		UserID       string                 `json:"user_id" bson:"_user_id"`
		TotalBalance float64                `bson:"total_balance" json:"total_balance"`
		Transactions []TransactionsPerMonth `bson:"transactions" json:"transactions"`
		TotalDebit   float64                `bson:"total_debit" json:"total_debit"`
		TotalCredit  float64                `bson:"total_credit" json:"total_credit"`
		CreatedAt    primitive.DateTime     `bson:"created_at,omitempty" json:"created_at,omitempty"`
		UpdatedAt    primitive.DateTime     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
		DeletedAt    primitive.DateTime     `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	}
)
