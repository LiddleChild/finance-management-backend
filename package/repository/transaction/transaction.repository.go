package transaction

import (
	"backend/core/models"
	"backend/core/repository/transaction"
)

type ITransactionRepository interface {
	GetTransactionsInRangeByUserId(userId string, startEpoch int64, endEpoch int64) ([]models.Transaction, error)
	CreateTransaction(userId string, transaction models.Transaction) error
	UpdateTransaction(userId string, transaction models.Transaction) error
	DeleteTransaction(userId string, transaction models.DeletingTransaction) error
}

func NewTransactionRepository() ITransactionRepository {
	return transaction.NewTransactionRepository()
}
