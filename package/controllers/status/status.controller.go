package status

import (
	"backend/core/controllers/status"

	"github.com/gofiber/fiber/v2"
)

type IStatusController interface {
	Ping(c *fiber.Ctx) error
}

func NewStatusController() IStatusController {
	return status.NewStatusController()
}
