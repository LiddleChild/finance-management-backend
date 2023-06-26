package models

type Transaction struct {
	TransactionId	string		`json:"TransactionId"`
	Amount				float64		`json:"Amount"`
	Note					string		`json:"Note"`
	Timestamp			int64			`json:"Timestamp"`
	Category			Category	`json:"Category"`
	Wallet				Wallet		`json:"Wallet"`
}