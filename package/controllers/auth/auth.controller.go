package auth

import (
	"backend/core/controllers/auth"
	"backend/package/repository/user"

	"github.com/gofiber/fiber/v2"
)

type IAuthController interface {
	CreateUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

func NewAuthController(repo user.IUserRepository) IAuthController {
	return auth.NewAuthController(repo)
}
