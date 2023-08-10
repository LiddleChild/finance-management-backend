package category

import (
	control "backend/core/controllers/category"
	repo "backend/package/repository/category"

	"github.com/gofiber/fiber/v2"
)

type ICategoryController interface {
	GetCategoryMap(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}

func NewCategoryController(repo repo.ICategoryRepository) ICategoryController {
	return control.NewCategoryController(repo)
}
