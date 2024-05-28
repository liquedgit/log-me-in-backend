package service

import (
	"log-me-in/model"
	"log-me-in/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AuthLogin(ctx *fiber.Ctx, username string, password []byte) error {
	user, err := GetUserByUsername(username)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponseDTO{
			Error: "Incorrect username or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), password); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponseDTO{
			Error: "Incorrect username or password",
		})
	}

	strToken, err := utils.GenerateNewToken(user.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponseDTO{
			Error: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.PostLoginResponseDTO{
		Message: "Logged in successfully",
		Token:   *strToken,
	})

}
