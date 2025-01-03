package handlers

import (
	"Banking/config"
	"Banking/models"

	"github.com/gofiber/fiber/v2"
)

func GetSaving(c *fiber.Ctx) error {
	clientID, ok := c.Locals("client_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized or invalid client ID",
		})
	}
	var savings []models.Savings

	if err := config.DB.Where("ClientID = ?", clientID).Find(&savings).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error retrieving savings",
		})
	}
	return c.Status(fiber.StatusOK).JSON(savings)

}
