package models

type Transaction struct {
	TransactionId string  `json:"TransactionId"`
	Amount        float64 `json:"Amount"`
	Note          string  `json:"Note"`
	Timestamp     int64   `json:"Timestamp"`
	Category      string  `json:"Category"`
	Wallet        string  `json:"Wallet"`
}

type CreatingTransaction struct {
	Amount    float64 `json:"Amount"    validate:"required"`
	Note      string  `json:"Note"`
	Timestamp int64   `json:"Timestamp" validate:"required"`
	Category  string  `json:"Category"  validate:"required"`
	Wallet    string  `json:"Wallet"    validate:"required"`
}

type TransactionFilter struct {
	Month int `query:"month"`
	Year  int `query:"year"`
}
