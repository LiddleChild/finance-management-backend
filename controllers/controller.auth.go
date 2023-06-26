package controllers

import (
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)


func CreateUser(c *fiber.Ctx) error {
	registeringUser := models.User{}
	err := c.BodyParser(&registeringUser)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.ErrorMessage("Empty user credentials"))
	}
	
	return nil
}