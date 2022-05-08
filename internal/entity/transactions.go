package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Transaction struct {
		ID          int    `json:"id" bson:"id"`
		Date        string `bson:"date" json:"date"`
		Transaction string `bson:"transaction" json:"transaction"`
	}
	TransactionsPerMonth struct {
		Month string
		Total int
	}
	TransactionHistory struct {
		ID           string                 `json:"id" bson:"_id"`
		Balance      float64                `json:"total_balance" bson:"total_balance"`
		Transactions []TransactionsPerMonth `json:"transactions_per_month" bson:"transactions_per_month"`
		Debit        float64                `json:"average_debit_amount" bson:"average_debit_amount"`
		Credit       float64                `json:"average_credit_amount" bson:"average_credit_amount"`
		CreatedAt    primitive.DateTime     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	}
)
