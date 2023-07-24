package databases

import (
	"backend/models"
	"backend/utils"
	"context"

	"google.golang.org/api/iterator"
)

func GetCategoryMapByUserId(userId string) (map[string]models.Category, error) {
	dbClient := utils.GetFirestoreClient()
	ctx := context.Background()
	
	categoryMap := map[string]models.Category{}

	itr := dbClient.Collection("user").Doc(userId).Collection("category").Documents(ctx)
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

	_, err := dbClient.Collection("user").Doc(userId).Collection("category").Doc(categoryId).Get(ctx)
	
	return err == nil
}