package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimMichel/ecommerce-api/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Kim!")
	})

	app.Listen(":3000")
}
