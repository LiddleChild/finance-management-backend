package auth

import (
	"github.com/gofiber/fiber/v2"
)

type AuthMockMiddleware struct{}

func NewMock() *AuthMockMiddleware {
	return &AuthMockMiddleware{}
}

func (mw *AuthMockMiddleware) RequireAccessToken(handler func(*fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Locals("UserId", "fmWEwAx6QXokS5xnICKW")

		return handler(c)
	}
}
