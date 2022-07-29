package main

import (
	"dumbmerch-api/database"
	"dumbmerch-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.DatabaseInit()
	database.Migration()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	// INITIAL ROUTE
	routes.RouteInit(api)

	app.Listen("localhost:5000")
}
