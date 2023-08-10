package auth

import (
	"backend/core/models"
	"backend/package/repository/user"
	"backend/utils"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	repo user.IUserRepository
}

func NewAuthController(repo user.IUserRepository) *AuthController {
	return &AuthController{
		repo,
	}
}

func (con *AuthController) CreateUser(c *fiber.Ctx) error {
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
	if con.repo.DoesUserExistByField("Email", registeringUser.Email) {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("User already exists"))
	}

	// Hash password
	utils.SaltAndHashPassword(&registeringUser.Password)

	// Create user in db
	err = con.repo.CreateUser(registeringUser)
	if err != nil {
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't create user"))
	}

	return c.
		Status(http.StatusCreated).
		SendString(utils.JSONMessage("User created successfully"))
}

func (con *AuthController) Login(c *fiber.Ctx) error {
	// Get body from request
	userCredentials := models.UserCredentials{}
	err := c.BodyParser(&userCredentials)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage("Empty user credentials"))
	}

	// Validate user information
	err = utils.GetValidator().Struct(userCredentials)
	if err != nil {
		errs := utils.ErrorsToString(utils.TranslateError(err))
		return c.
			Status(http.StatusBadRequest).
			SendString(utils.JSONMessage(
				fmt.Sprintf("Invalid user credentials: %s", strings.Join(errs, ", "))))
	}

	// Get user from email
	user, err, ok := con.repo.GetUserByField("Email", userCredentials.Email)
	if err != nil {
		fmt.Println(err.Error())
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't get user credentials"))
	}

	// Check for valid email with matched password
	if !ok || !utils.CheckPassword(user.Password, userCredentials.Password) {
		return c.
			Status(http.StatusUnauthorized).
			SendString(utils.JSONMessage("Email and Password mismatch"))
	}

	// JWT expired in 1 days
	expireDate := time.Now().Add(time.Hour * 24)

	// Create claim
	claim := models.JWTClaim{
		UserId: user.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireDate),
		},
	}

	// Create JWT token with claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Sign token with secret key
	accessToken, err := token.SignedString([]byte(utils.GetEnv("JWT_PRIVATE_KEY", "")))
	if err != nil {
		fmt.Println(err.Error())
		return c.
			Status(http.StatusConflict).
			SendString(utils.JSONMessage("Couldn't log you in"))
	}

	// Create httpOnly cookie
	cookie := &fiber.Cookie{
		HTTPOnly: true,
		Name:     "access_token",
		Value:    accessToken,
		Expires:  expireDate,
	}

	// Set cookie
	c.Cookie(cookie)

	return c.
		Status(http.StatusOK).
		SendString(utils.JSONMessage("You are logged in!"))
}

func (con *AuthController) Logout(c *fiber.Ctx) error {
	// Create empty cookie
	cookie := &fiber.Cookie{
		HTTPOnly: true,
		Name:     "access_token",
		Value:    "",
	}

	// Set cookie
	c.Cookie(cookie)

	return c.
		Status(http.StatusOK).
		SendString(utils.JSONMessage("You are logged out"))
}
