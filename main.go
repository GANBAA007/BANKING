package main

import (
	"Banking/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not loading")
	}
	config.ConnectDB()
	config.Migrate()
}
