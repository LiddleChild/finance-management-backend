package transaction

import (
	"backend/core/models"
	"backend/package/repository/category"
	"backend/package/repository/transaction"
	"backend/package/repository/wallet"
	"backend/utils"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	transactionRepo transaction.ITransactionRepository
	walletRepo      wallet.IWalletRepository
	categoryRepo    category.ICategoryRepository
}

func NewTransactionController(
	transactionRepo transaction.ITransactionRepository,
	walletRepo wallet.IWalletRepository,
	categoryRepo category.ICategoryRepository) *TransactionController {

	return &TransactionController{
		transactionRepo,
		walletRepo,
		categoryRepo,
	}
}

func (con *TransactionController) GetTransaction(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)
	filter := models.TransactionFilter{Month: -1, Year: -1}
	c.QueryParser(&filter)

	var startEpoch int64 = math.MinInt64
	var endEpoch int64 = math.MaxInt64
	if filter.Month > 0 && filter.Year > 0 {
		startEpoch, endEpoch = utils.GetMonthsRange(filter.Month, filter.Year, 1)
	}

	transactionLists, err := con.transactionRepo.GetTransactionsInRangeByUserId(userId, startEpoch, endEpoch)
	if err != nil {
		fmt.Println(err)
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't get transactions"))
	}

	jsonStr, err := json.Marshal(transactionLists)

	return c.
		Status(http.StatusOK).
		SendString(string(jsonStr))
}

func (con *TransactionController) GetTodayTransaction(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)

	startEpoch, endEpoch := utils.GetTodayRange()
	transactionLists, err := con.transactionRepo.GetTransactionsInRangeByUserId(userId, startEpoch, endEpoch)
	if err != nil {
		fmt.Println(err)
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't get transactions"))
	}

	jsonStr, err := json.Marshal(transactionLists)

	return c.
		Status(http.StatusOK).
		SendString(string(jsonStr))
}

func (con *TransactionController) CreateTransaction(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)

	transaction := models.Transaction{
		TransactionId: utils.GenerateUUID(),
	}

	// Parse body
	err := c.BodyParser(&transaction)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty transaction information"))
	}

	// Validate user information
	err = utils.GetValidator().Struct(transaction)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid transaction information: %s", strings.Join(errs, ", "))))
	}

	// Validate wallet and category ids
	if !(con.walletRepo.DoesWalletExist(userId, transaction.Wallet) &&
		con.categoryRepo.DoesCategoryExist(userId, transaction.Category)) {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Invalid wallet or category"))
	}

	// Create transaction
	err = con.transactionRepo.CreateTransaction(userId, transaction)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't create transaction"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("Transaction created"))
}

/*
Update transaction
*/
func (con *TransactionController) UpdateTransaction(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)

	// Parse body
	transaction := models.Transaction{}
	err := c.BodyParser(&transaction)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty transaction information"))
	}

	// Validate user information
	err = utils.GetValidator().Struct(transaction)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid transaction information: %s", strings.Join(errs, ", "))))
	}

	// Update transaction
	err = con.transactionRepo.UpdateTransaction(userId, transaction)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't update transaction"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("Transaction updated"))
}

/*
Delete transaction
*/
func (con *TransactionController) DeleteTransaction(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)

	// Parse body
	deletingTransaction := models.DeletingTransaction{}
	err := c.BodyParser(&deletingTransaction)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty transaction information"))
	}

	// Validate user information
	err = utils.GetValidator().Struct(deletingTransaction)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid transaction information: %s", strings.Join(errs, ", "))))
	}

	// Delete category
	err = con.transactionRepo.DeleteTransaction(userId, deletingTransaction)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't delete transaction"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("Transaction deleted"))
}
