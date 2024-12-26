package handlers

import (
	"Banking/config"
	"Banking/models"

	"github.com/gofiber/fiber/v2"
)

func GetAcc(c *fiber.Ctx) error {
	clientID, ok := c.Locals("client_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized or invalid client ID",
		})
	}

	var accounts []models.Wallet

	if err := config.DB.Where("ClientID = ?", clientID).Find(&accounts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error retrieving accounts",
		})
	}

	return c.Status(fiber.StatusOK).JSON(accounts)
}
