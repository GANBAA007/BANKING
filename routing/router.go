package routing

import (
	handlers "Banking/Handlers"
	accounthandler "Banking/accountHandler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/login", handlers.CheckLogin)
	auth.Post("/createClient", handlers.CreateClient)
	user := app.Group("/user")
	user.Post("/getacc", handlers.GetAccAll)
	user.Post("/getsave", handlers.GetSaving)
	user.Post("/transaction", accounthandler.Transferfunds)
	user.Post("/createsaving", accounthandler.CreateSaving)
	user.Post("/createacc", accounthandler.CreateAccount)
}
