package controllers

import (
	"time"

	"github.com/ItzRyz/squareville-api/src/database"
	"github.com/ItzRyz/squareville-api/src/models"
	"github.com/ItzRyz/squareville-api/src/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(models.Login)

	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to Login",
			"error":   errValidate.Error(),
		})
	}

	var user models.User

	errEmail := database.DB.Debug().First(&user, "username = ?", loginRequest.Username).Error
	if errEmail != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	isValid := utils.CheckHash(loginRequest.Password, user.Password)
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password is not valid",
		})
	}

	claims := jwt.MapClaims{}
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(2 * time.Minute).Unix()

	token, errGenerate := utils.GenerateToken(&claims)

	if errGenerate != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Token is not valid",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
