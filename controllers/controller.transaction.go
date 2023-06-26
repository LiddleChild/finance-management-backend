package controllers

import (
	"backend/databases"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"

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