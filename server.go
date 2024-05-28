package main

import (
	"fmt"
	"log-me-in/database"
	"log-me-in/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			// Handle the panic or return an error response
		}
	}()
	app := fiber.New()

	database.AutoMigrate()
	router.SetupRoutes(app)

	err := app.Listen("127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
}
