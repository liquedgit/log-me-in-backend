package service

import (
	"log-me-in/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AuthLogin(ctx *fiber.Ctx, username string, password []byte) error {
	user, err := GetUserByUsername(username)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect username or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), password); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect username or password",
		})
	}

	strToken, err := utils.GenerateNewToken(user.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":   strToken,
		"message": "Logged in succesfully",
	})

}
