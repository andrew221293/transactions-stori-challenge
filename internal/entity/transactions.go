package entity

type (
	Transaction struct {
		ID          int    `json:"id" bson:"id"`
		Date        string `bson:"date" json:"date"`
		Transaction string `bson:"transaction" json:"transaction"`
	}
)
