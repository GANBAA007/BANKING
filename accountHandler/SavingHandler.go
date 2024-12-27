package accounthandler

import (
	"Banking/models"

	"github.com/gofiber/fiber/v2"
)

func CreateSaving(c *fiber.Ctx) error {
	var saving []models.Savings

	if err := c.BodyParser(&saving); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(saving)
}
