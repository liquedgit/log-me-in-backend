package router

import (
	"fmt"
	"log-me-in/router/routes"
	"log-me-in/utils"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/check-jwt", func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		token = token[7:]
		jwtObj, err := utils.DecodeToken(token)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		fmt.Println(jwtObj.Claims["id"])
		return ctx.SendString(token)
	})
	routes.SetupAuthRoutes(api)
	routes.SetupUserRoutes(api)
	routes.SetupNoteRoutes(api)
}
