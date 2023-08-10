package category

import (
	"backend/core/models"
	"backend/utils"
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type CategoryRepository struct{}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (repo *CategoryRepository) GetCategoryMapByUserId(userId string) (map[string]models.Category, error) {
	db := utils.GetFirestoreClient()
	ctx := context.Background()

	categoryMap := map[string]models.Category{}

	itr := db.Collection("user").
		Doc(userId).
		Collection("category").
		Documents(ctx)

	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return categoryMap, err
		}

		category := models.Category{}
		doc.DataTo(&category)

		categoryMap[category.CategoryId] = category
	}

	return categoryMap, nil
}

func (repo *CategoryRepository) DoesCategoryExist(userId string, categoryId string) bool {
	db := utils.GetFirestoreClient()
	ctx := context.Background()

	_, err := db.Collection("user").
		Doc(userId).
		Collection("category").
		Doc(categoryId).
		Get(ctx)

	return err == nil
}

func (repo *CategoryRepository) IsCategoryEditable(userId string, categoryId string) (bool, error) {
	db := utils.GetFirestoreClient()
	ctx := context.Background()

	doc, err := db.Collection("user").
		Doc(userId).
		Collection("category").
		Doc(categoryId).
		Get(ctx)
	if err != nil {
		return false, err
	}

	category := models.Category{}
	err = doc.DataTo(&category)
	if err != nil {
		return false, err
	}

	return category.Editable, nil
}

func (repo *CategoryRepository) CreateCategory(userId string, category models.Category) error {
	db := utils.GetFirestoreClient()
	ctx := context.Background()

	_, err := db.Collection("user").
		Doc(userId).
		Collection("category").
		Doc(category.CategoryId).
		Set(ctx, category)

	return err
}

func (repo *CategoryRepository) UpdateCategory(userId string, category models.Category) error {
	db := utils.GetFirestoreClient()
	ctx := context.Background()

	_, err := db.Collection("user").
		Doc(userId).
		Collection("category").
		Doc(category.CategoryId).
		Update(ctx, []firestore.Update{
			{
				Path:  "Label",
				Value: category.Label,
			}, {
				Path:  "Color",
				Value: category.Color,
			},
		})

	return err
}

func (repo *CategoryRepository) DeleteCategory(userId string, category models.DeletingCategory) error {
	db := utils.GetFirestoreClient()
	ctx := context.Background()

	doc, err := db.Collection("user").
		Doc(userId).
		Collection("category").
		Doc(category.CategoryId).
		Get(ctx)
	if err != nil {
		return err
	}

	_, err = doc.Ref.Delete(ctx)
	return err
}
