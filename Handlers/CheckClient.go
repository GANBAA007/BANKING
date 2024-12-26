package handlers

import (
	"Banking/config"
	"Banking/models"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateClient(c *fiber.Ctx) error {
	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}

	var existingClient models.Client
	err := config.DB.Where("regno = ?", client.RegNo).First(&existingClient).Error //reg_existing
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Database query error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Database error",
			})
		}
	} else {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Registery number already exists.",
		})
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(client.Password), bcrypt.DefaultCost) //generate password
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not hash password"})
	}
	client.Password = string(hashedpassword)

	if err := config.DB.Create(&client).Error; err != nil {
		log.Printf("Failed to save client: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not save client",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(client)
}
