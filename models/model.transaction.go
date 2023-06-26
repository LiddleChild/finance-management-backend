package models

type Transaction struct {
	TransactionId	string	`json:"TransactionId"`
	Amount				float64	`json:"Amount"`
	Note					string	`json:"Note"`
	Timestamp			int64		`json:"Timestamp"`
	Category			string	`json:"Category"`
	Wallet				string	`json:"Wallet"`
}