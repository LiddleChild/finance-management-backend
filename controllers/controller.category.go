package controllers

import (
	"backend/databases"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetCategoryMap(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)
	categoryMap, err := databases.GetCategoryMapByUserId(userId)
	if err != nil {
		if err != nil {
			fmt.Println(err)
			return c.
				Status(http.StatusConflict).
				SendString(utils.JSONMessage("Couldn't get categories"))
		}
	}

	jsonStr, err := json.Marshal(categoryMap)

	return c.
		Status(http.StatusOK).
		SendString(string(jsonStr))
}