package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimMichel/ecommerce-api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListProducts)

	app.Post("/products", handlers.CreateProducts)
}
