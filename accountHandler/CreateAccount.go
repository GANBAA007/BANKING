package accounthandler

import (
	"Banking/config"
	"Banking/models"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateAccount(c *fiber.Ctx) error {
	var account models.Wallet

	if err := c.BodyParser(&account); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	var existingacc models.Wallet
	err := config.DB.Where("account_no=?", existingacc.AccountNo).First(&existingacc).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Database query error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Database error",
			})
		}
	} else {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Дансны дугаар дөвхцаж байна",
		})
	}
	if err := config.DB.Create(&account).Error; err != nil {
		log.Printf("Failed to save client: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not save client",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(account)

}
