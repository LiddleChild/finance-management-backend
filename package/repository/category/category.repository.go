package category

import (
	"backend/core/models"
	"backend/core/repository/category"
)

type ICategoryRepository interface {
	GetCategoryMapByUserId(userId string) (map[string]models.Category, error)
	DoesCategoryExist(userId string, categoryId string) bool
	IsCategoryEditable(userId string, categoryId string) (bool, error)
	CreateCategory(userId string, category models.Category) error
	UpdateCategory(userId string, category models.Category) error
	DeleteCategory(userId string, category models.DeletingCategory) error
}

func NewCategoryRepository() ICategoryRepository {
	return category.NewCategoryRepository()
}
