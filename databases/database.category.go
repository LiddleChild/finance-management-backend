package databases

import (
	"backend/models"
	"backend/utils"
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetCategoryMapByUserId(userId string) (map[string]models.Category, error) {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	categoryMap := map[string]models.Category{}

	itr := dbClient.Collection("user").
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

		// Set id
		category.CategoryId = doc.Ref.ID

		categoryMap[category.CategoryId] = category
	}

	return categoryMap, nil
}

func DoesCategoryExist(userId string, categoryId string) bool {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	_, err := dbClient.Collection("user").
		Doc(userId).
		Collection("category").
		Doc(categoryId).
		Get(ctx)

	return err == nil
}

func IsCategoryEditable(userId string, categoryId string) (bool, error) {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	doc, err := dbClient.Collection("user").
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

func CreateCategory(userId string, category models.Category) error {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	_, err := dbClient.Collection("user").
		Doc(userId).
		Collection("category").
		Doc(category.CategoryId).
		Set(ctx, category)

	return err
}

func UpdateCategory(userId string, category models.Category) error {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	_, err := dbClient.Collection("user").
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

func DeleteCategory(userId string, category models.DeletingCategory) error {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()

	doc, err := dbClient.Collection("user").
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
