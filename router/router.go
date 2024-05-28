package router

import (
	"log-me-in/middleware"
	"log-me-in/model"
	"log-me-in/router/routes"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/uploads/*", middleware.TokenMiddleware, func(ctx *fiber.Ctx) error {
		filePath, err := url.QueryUnescape(ctx.Params("*"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponseDTO{
				Error: "You have sended a bad request",
			})
		}
		cwd, err := os.Getwd()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponseDTO{
				Error: "Somethings wrong with the server",
			})
		}
		finalPath := filepath.Join(cwd, "uploads", filePath)
		if _, err := os.Stat(finalPath); err != nil {
			if os.IsNotExist(err) {
				// File not found
				return ctx.Status(fiber.StatusNotFound).SendString("File not found")
			}
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return ctx.SendFile(finalPath)
	})
	api := app.Group("/api")

	routes.SetupAuthRoutes(api)
	routes.SetupUserRoutes(api)
	routes.SetupNoteRoutes(api)
}
