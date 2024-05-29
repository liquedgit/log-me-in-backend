package main

import (
	"log-me-in/database"
	"log-me-in/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	database.AutoMigrate()
	router.SetupRoutes(app)

	err := app.Listen("0.0.0.0:1234")
	if err != nil {
		panic(err)
	}
}
