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
		Reg_no   string `json:"regno"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&Login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	Login.Reg_no = strings.TrimSpace(Login.Reg_no)

	var clients models.Client
	if err := config.DB.Where("reg_no=?", Login.Reg_no).First(&clients).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Email or password incorrect"})
		}
		log.Printf("Error fetching user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An unexpected error occurred",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(clients.Password), []byte(Login.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Email or password incorrect",
		})

	}

	token, err := token.GenerateToken(clients.ID)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"ClientID": clients.ID,
			"regno":    clients.RegNo,
			"name":     clients.Name,
			"token":    token,
		},
	})
}
