package routing

import (
	handlers "Banking/Handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/login", handlers.CheckLogin)
	auth.Post("/createClient", handlers.CreateClient)
}
