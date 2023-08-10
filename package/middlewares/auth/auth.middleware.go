package auth

import (
	"github.com/gofiber/fiber/v2"
)

type IAuthMiddleware interface {
	RequireAccessToken(handler func(*fiber.Ctx) error) func(c *fiber.Ctx) error
}

func NewMiddleware(auth IAuthMiddleware) IAuthMiddleware {
	return auth
}
