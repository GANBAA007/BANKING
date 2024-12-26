package handlers

import (
	token "Banking/Token"
	"Banking/config"
	"Banking/models"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckLogin(c *fiber.Ctx) error {
	var Login struct {
		RegNo    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&Login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	Login.RegNo = strings.TrimSpace(Login.RegNo)

	var client models.Client
	if err := config.DB.Where("regno=?", Login.RegNo).First(&client).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Email or password incorrect"})
		}
		log.Printf("Error fetching user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An unexpected error occurred",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(Login.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Email or password incorrect",
		})

	}

	token, err := token.GenerateToken(client.ID)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"ClientID": client.ID,
			"regno":    client.RegNo,
			"name":     client.Name,
			"token":    token,
		},
	})
}
