package entity

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
)
