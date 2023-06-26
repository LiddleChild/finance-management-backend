package models

type Transaction struct {
	TransactionId	string
	Amount				float64
	Note					string
	Timestamp			int64
	Category			Category
	Wallet				Wallet
}