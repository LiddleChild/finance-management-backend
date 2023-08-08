package controllers

import (
	"backend/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.
		Status(http.StatusOK).
		SendString(utils.JSONMessage("Pong"))
}
