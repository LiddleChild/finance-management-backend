package controllers

import (
	"backend/databases"
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetTransaction(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)
	transactionLists, err := databases.GetTransactionsByUserId(userId)
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

func CreateTransaction(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)

	// Parse body
	creatingTransaction := models.CreatingTransaction{}
	err := c.BodyParser(&creatingTransaction)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty transaction information"))
	}

	// Validate user information
	err = utils.GetValidator().Struct(creatingTransaction)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid transaction information: %s", strings.Join(errs, ", "))))
	}

	// Validate wallet and category ids
	if !(databases.DoesWalletExist(userId, creatingTransaction.Wallet) &&
		databases.DoesCategoryExist(userId, creatingTransaction.Category)) {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Invalid wallet or category"))
	}

	// Create transaction
	err = databases.CreateTransaction(userId, creatingTransaction)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't create transaction"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("Transaction created"))
}