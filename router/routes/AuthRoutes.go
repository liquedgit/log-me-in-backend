package routes

import (
	"log-me-in/model"
	"log-me-in/service"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router) {
	api.Post("/login", func(ctx *fiber.Ctx) error {
		loginDTO := new(model.LoginRequestDTO)
		if err := ctx.BodyParser(loginDTO); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponseDTO{
				Error: err.Error(),
			})
		}

		if len(loginDTO.Username) <= 0 || len(loginDTO.Password) <= 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponseDTO{
				Error: "Username or password cannot be empty",
			})
		}

		return service.AuthLogin(ctx, loginDTO.Username, []byte(loginDTO.Password))
	})

	api.Post("/register", func(ctx *fiber.Ctx) error {
		requestDTO := new(model.RegisterRequestDTO)
		if err := ctx.BodyParser(requestDTO); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponseDTO{
				Error: err.Error(),
			})
		}

		if len(requestDTO.Username) <= 0 || len(requestDTO.Password) <= 0 || len(requestDTO.ConfirmPassword) <= 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponseDTO{
				Error: "Username , password or confirm password cannot be empty",
			})
		}

		if requestDTO.Password != requestDTO.ConfirmPassword {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponseDTO{
				Error: "Password and confirm password must be the same",
			})
		}

		_, err := service.GetUserByUsername(requestDTO.Username)

		if err == nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponseDTO{
				Error: "Username must be unique",
			})
		}

		user, err := service.CreateNewUser(requestDTO.Username, requestDTO.Password)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponseDTO{
				Error: "Somethings wrong with the server",
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(model.RegisterResponseDTO{
			Username: user.Username,
			Message:  "Successfully created new account",
		})
	})
}
