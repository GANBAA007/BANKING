package main

import (
	"Banking/config"
	"Banking/routing"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not loading")
	}
	config.ConnectDB()
	config.Migrate()
	app := fiber.New()
	routing.SetupRoutes(app)
	log.Println("Starting the server...")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	log.Println("Server stopped unexpectedly")

}
