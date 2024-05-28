package routes

import (
	"log-me-in/middleware"
	"log-me-in/model"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router) {
	user := api.Group("/user")
	user.Get("/", middleware.TokenMiddleware, func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*model.User)
		return ctx.Status(fiber.StatusOK).JSON(model.GetUserResponseDTO{
			Id:       user.Id,
			Username: user.Username,
			Role:     user.Role,
		})
	})
}
