package status

import (
	"backend/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type StatusController struct{}

func NewStatusController() *StatusController {
	return &StatusController{}
}

func (con *StatusController) Ping(c *fiber.Ctx) error {
	return c.
		Status(http.StatusOK).
		SendString(utils.JSONMessage("Pong"))
}
