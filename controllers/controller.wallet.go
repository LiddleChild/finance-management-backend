package controllers

import (
	"backend/databases"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetWalletMap(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)
	walletMap, err := databases.GetWalletMapByUserId(userId)
	if err != nil {
		if err != nil {
			fmt.Println(err)
			return c.
				Status(http.StatusConflict).
				SendString(utils.JSONMessage("Couldn't get wallets"))
		}
	}

	jsonStr, err := json.Marshal(walletMap)

	return c.
		Status(http.StatusOK).
		SendString(string(jsonStr))
}