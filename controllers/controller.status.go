package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
