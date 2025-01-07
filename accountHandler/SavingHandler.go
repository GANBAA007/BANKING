package accounthandler

import (
	"Banking/config"
	"Banking/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type req struct {
	InitAmount     float64 `json:"initAmount"`
	Duration       int     `json:"duration"`
	IsChecking     bool    `json:"Checking"`
	MonthlyPayment float64 `json:"monthlypayment"`
}

func CreateSaving(c *fiber.Ctx) error {
	var request req
	sixMonths := 6
	oneYear := 12
	twoYears := 24

	var saving models.Savings

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	if request.InitAmount < 1 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Must be greater than 1",
		})
	}

	if request.MonthlyPayment < 20000 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Monthly payment Must be greater than 20000",
		})
	}
	clientID, ok := c.Locals("client_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized or invalid client ID",
		})
	}

	if request.Duration < sixMonths {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Duration must be at least six months",
		})
	} else if request.Duration < oneYear {
		saving.InterestRate = 0.005
	} else if request.Duration < twoYears {
		saving.InterestRate = 0.01
	} else {
		saving.InterestRate = 0.05
	}
	saving.ClientID = clientID
	saving.MinAmount = request.InitAmount
	saving.ExpDate = time.Now().AddDate(0, request.Duration, 0)
	saving.CreatedAt = time.Now()
	saving.UpdatedAt = time.Now()
	saving.MonthlyPayment = request.MonthlyPayment
	saving.PaymentDate = time.Now()
	saving.Balance = request.InitAmount
	saving.InterestProfit = 0.00

	if err := config.DB.Create(&saving).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create savings record",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(saving)

}

func Calcinterest(c *fiber.Ctx) error {
	var saving models.Savings
	if err := config.DB.Where("ID=?", saving.ID).First(&saving).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Saving acc not found"})
		}
		log.Printf("Error fetching user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An unexpected error occurred",
		})
	}
	profit := saving.Balance * saving.InterestRate

	saving.InterestProfit = profit + saving.InterestProfit

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      saving.ID,
		"Balance": saving.Balance,
	})
}

func Cancelsaving(c *fiber.Ctx) error {
	var req struct {
		savingID uint
	}
	var saving models.Savings
	var Wallet models.Wallet
	if err := config.DB.Where("ID=?", req.savingID).First(&saving).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Saving acc not found"})
		}
		log.Printf("Error fetching user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An unexpected error occurred",
		})
	}
	clientID, ok := c.Locals("client_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized or invalid client ID",
		})
	}
	if err := config.DB.Where("ClientID=?", clientID).First(&Wallet).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "wallet not found",
		})
	}

	var transferReq Req
	transferReq.SrcAcc = saving.AccountNo
	transferReq.DestAcc = Wallet.AccountNo
	transferReq.Amount = saving.Balance

	if err := Transferfunds(c); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to transfer funds",
		})
	}

	if err := config.DB.Where("ID=?", req.savingID).Delete(&saving).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Already canceled"})
		}
		log.Printf("Error fetching user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An unexpected error occurred",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      saving.ID,
		"message": "Success saving account canceled",
	})
}
