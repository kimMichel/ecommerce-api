package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimMichel/ecommerce-api/database"
	"github.com/kimMichel/ecommerce-api/models"
)

func ListProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.DB.Db.Find(&products)

	return c.Status(200).JSON(products)
}

func CreateProducts(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&product)

	return c.Status(200).JSON(product)
}
