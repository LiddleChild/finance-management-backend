package middlewares

import (
	"backend/models"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAccessToken(handler func(*fiber.Ctx) error) func(*fiber.Ctx) error {
	return func (c *fiber.Ctx) error {
		accessToken := c.Cookies("access_token", "")

		// Check for access token
		if accessToken == "" {
			return c.
				Status(http.StatusUnauthorized).
				SendString("Login required")
		}
		
		// Validate access token
		claim := &models.JWTClaim{}
		_, err := jwt.ParseWithClaims(accessToken, claim,
			func (t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_PRIVATE_KEY")), nil
			})

		if err != nil {
			return c.
				Status(http.StatusUnauthorized).
				SendString("Please login to continue")
		}

		// Send user id to handler
		c.Locals("UserId", claim.UserId)

		return handler(c)
	}
}