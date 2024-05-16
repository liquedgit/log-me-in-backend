package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log-me-in/model"
	"log-me-in/router/routes"
	"log-me-in/service"
	"log-me-in/utils"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/login", func(ctx *fiber.Ctx) error {
		loginDTO := new(model.LoginRequestDTO)
		if err := ctx.BodyParser(loginDTO); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if len(loginDTO.Username) <= 0 || len(loginDTO.Password) <= 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Username or password cannot be empty",
			})
		}

		return service.AuthLogin(ctx, loginDTO.Username, []byte(loginDTO.Password))
	})

	api.Post("/register", func(ctx *fiber.Ctx) error {
		requestDTO := new(model.RegisterRequestDTO)
		if err := ctx.BodyParser(requestDTO); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if len(requestDTO.Username) <= 0 || len(requestDTO.Password) <= 0 || len(requestDTO.ConfirmPassword) <= 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Username , password or confirm password cannot be empty",
			})
		}

		if requestDTO.Password == requestDTO.ConfirmPassword {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Password and confirm password must be the same",
			})
		}

		user, err := service.CreateNewUser(requestDTO.Username, requestDTO.Password)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Somethings wrong with the server",
			})
		}
		responseDTO := model.RegisterResponseDTO{
			Username: user.Username,
			Message:  "Successfully created new account",
		}
		return ctx.Status(fiber.StatusCreated).JSON(responseDTO)
	})
	api.Get("/check-jwt", func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		token = token[7:]
		jwtObj, err := utils.DecodeToken(token)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		fmt.Println(jwtObj.Claims["role"])
		return ctx.SendString(token)
	})
	routes.SetupUserRoutes(api)
}
