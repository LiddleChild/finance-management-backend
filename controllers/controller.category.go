package controllers

import (
	"backend/databases"
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

/*
Get category
*/
func GetCategoryMap(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)
	categoryMap, err := databases.GetCategoryMapByUserId(userId)
	if err != nil {
		if err != nil {
			fmt.Println(err)
			return c.
				Status(http.StatusConflict).
				SendString(utils.JSONMessage("Couldn't get categories"))
		}
	}

	jsonStr, err := json.Marshal(categoryMap)

	return c.
		Status(http.StatusOK).
		SendString(string(jsonStr))
}

/*
Create category
*/
func CreateCategory(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)

	category := models.Category{
		CategoryId: utils.GenerateUUID(),
		Editable:   true,
	}

	// Parse body
	err := c.BodyParser(&category)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty category information"))
	}

	// Validate user information
	err = utils.GetValidator().Struct(category)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid category information: %s", strings.Join(errs, ", "))))
	}

	// Create category
	err = databases.CreateCategory(userId, category)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't create category"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("Category created"))
}

/*
Update category
*/
func UpdateCategory(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)

	// Parse body
	category := models.Category{}
	err := c.BodyParser(&category)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty transaction information"))
	}

	// Check if editable
	ok, err := databases.IsCategoryEditable(userId, category.CategoryId)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Cannot find category"))
	} else if !ok {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Cannot edit category"))
	}

	// Validate user information
	err = utils.GetValidator().Struct(category)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid category information: %s", strings.Join(errs, ", "))))
	}

	// Update category
	err = databases.UpdateCategory(userId, category)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't update category"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("Category updated"))
}

/*
Delete category
*/
func DeleteCategory(c *fiber.Ctx) error {
	userId := c.Locals("UserId").(string)

	// Parse body
	deletingCategory := models.DeletingCategory{}
	err := c.BodyParser(&deletingCategory)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty category information"))
	}

	// Validate user information
	err = utils.GetValidator().Struct(deletingCategory)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid category information: %s", strings.Join(errs, ", "))))
	}

	// Delete category
	err = databases.DeleteCategory(userId, deletingCategory)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't delete category"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("Category deleted"))
}
