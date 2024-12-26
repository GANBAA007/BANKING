package accounthandler

import "github.com/gofiber/fiber/v2"

type Req struct {
	SrcAcc  string  `json:"src"`
	DestAcc string  `json:"dest"`
	Amount  float64 `json:"amount"`
}

func Transferfunds(c *fiber.Ctx) error {

	var request Req

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
}
