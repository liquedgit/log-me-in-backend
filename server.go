package main

import (
	"github.com/gofiber/fiber/v2"
	"log-me-in/database"
	"log-me-in/router"
)

func main() {
	app := fiber.New()
	database.AutoMigrate()
	router.SetupRoutes(app)
	err := app.Listen("127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
}
