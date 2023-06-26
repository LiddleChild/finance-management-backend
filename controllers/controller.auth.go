package controllers

import (
	db "backend/databases"
	"backend/models"
	"backend/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)


func CreateUser(c *fiber.Ctx) error {
	// Parse body
	registeringUser := models.RegisteringUser{}
	err := c.BodyParser(&registeringUser)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty user credentials"))
	}

	// Validate body
	err = utils.GetValidator().Struct(registeringUser)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid user credentials: %s", strings.Join(errs, ", "))))
	}

	// Test for existing user
	if db.DoesUserExistByField("Email", registeringUser.Email) {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("User already exists"))
	}

	// Hash password
	utils.SaltAndHashPassword(&registeringUser.Password)

	// Create user in db
	err = db.CreateUser(registeringUser)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't create user"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("User created successfully"))
}