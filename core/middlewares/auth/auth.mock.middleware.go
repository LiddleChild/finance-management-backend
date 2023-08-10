package auth

import (
	"github.com/gofiber/fiber/v2"
)

type AuthMockMiddleware struct {
	userId string
}

func NewMock(userId string) *AuthMockMiddleware {
	return &AuthMockMiddleware{
		userId,
	}
}

func (mw *AuthMockMiddleware) RequireAccessToken(handler func(*fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Locals("UserId", mw.userId)

		return handler(c)
	}
}
