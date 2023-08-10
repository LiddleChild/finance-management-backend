package transaction

import (
	transactionControl "backend/core/controllers/transaction"
	"backend/package/repository/category"
	"backend/package/repository/transaction"
	"backend/package/repository/wallet"

	"github.com/gofiber/fiber/v2"
)

type ITransactionController interface {
	GetTransaction(c *fiber.Ctx) error
	GetTodayTransaction(c *fiber.Ctx) error
	CreateTransaction(c *fiber.Ctx) error
	UpdateTransaction(c *fiber.Ctx) error
	DeleteTransaction(c *fiber.Ctx) error
}

func NewTransactionController(
	transactionRepo transaction.ITransactionRepository,
	walletRepo wallet.IWalletRepository,
	categoryRepo category.ICategoryRepository) ITransactionController {

	return transactionControl.NewTransactionController(transactionRepo, walletRepo, categoryRepo)
}
