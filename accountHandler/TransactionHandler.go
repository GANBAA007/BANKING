package accounthandler

import (
	"Banking/config"
	"Banking/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Req struct {
	SrcAcc  string  `json:"src"`
	DestAcc string  `json:"dest"`
	Amount  float64 `json:"amount"`
}

func Transferfunds(c *fiber.Ctx) error {

	var request Req
	var srcAccount, destAccount models.Wallet

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	if request.DestAcc == request.SrcAcc {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Same accNo",
		})
	}
	if request.Amount < 1 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Must be greater than 1",
		})
	}

	if err := config.DB.Where("AccountNo=?", request.DestAcc).First(&destAccount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Dest Not valid AccountNo",
		})
	}
	if err := config.DB.Where("AccountNo=?", request.SrcAcc).First(&srcAccount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Src Not valid AccountNo",
		})
	}

	if strings.HasPrefix(request.DestAcc, "5") {
		if srcAccount.Balance < request.Amount+100 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Not enough funds",
			})
		}
	} else {
		if srcAccount.Balance < request.Amount+300 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Not enough funds",
			})
		}
	}

	// transaction starts

	srcAccount.Balance -= request.Amount
	destAccount.Balance += request.Amount
	var bankfee float64

	if strings.HasPrefix(request.DestAcc, "5") {
		bankfee = 100.00
		srcAccount.Balance -= bankfee

		var masterAccount models.Wallet
		if err := config.DB.Where("ID", 1).First(&masterAccount).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Master Account not found",
			})
		}
		masterAccount.Balance += bankfee

		if err := config.DB.Save(&srcAccount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update source account",
			})
		}

		if err := config.DB.Save(&masterAccount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update master account",
			})
		}

	} else {
		bankfee = 300.00
		srcAccount.Balance -= bankfee
		var masterAccount models.Wallet
		if err := config.DB.Where("ID", 1).First(&masterAccount).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Master Account not found",
			})
		}
		masterAccount.Balance += bankfee

		if err := config.DB.Save(&srcAccount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update source account",
			})
		}

		if err := config.DB.Save(&masterAccount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update master account",
			})
		}
	}

	if err := config.DB.Save(&srcAccount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update source account",
		})
	}

	if err := config.DB.Save(&destAccount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update destination account",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Transfer successful",
		"data": map[string]interface{}{
			"source_account":      srcAccount.AccountNo,
			"destination_account": destAccount.AccountNo,
			"amount_transferred":  request.Amount,
			"Bank_fee":            bankfee,
		},
	})

}
