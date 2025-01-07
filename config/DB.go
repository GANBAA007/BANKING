package config

import (
	"Banking/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	log.Println("Connecting to database with DSN: ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	DB = db

	return DB
}

func Migrate() {
	err := DB.AutoMigrate(&models.Client{}, &models.Wallet{}, &models.Savings{}, &models.Loan{})
	if err != nil {
		log.Fatal("Migrating failed:", err)
	} else {
		log.Println("migration success")
	}

}
