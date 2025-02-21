package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load the environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Debugging: Print environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	fmt.Println("DB_USER:", dbUser) // Debugging print
	fmt.Println("DB_PASS:", dbPass) // Debugging print
	fmt.Println("DB_NAME:", dbName) // Debugging print
	fmt.Println("DB_HOST:", dbHost) // Debugging print
	fmt.Println("DB_PORT:", dbPort) // Debugging print

	// Check if all required environment variables are present
	if dbUser == "" || dbPass == "" || dbName == "" || dbHost == "" || dbPort == "" {
		log.Fatal("Missing one or more database configuration values in .env")
	}

	// Build the DSN (Data Source Name)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbHost, dbPort)

	// Reuse the 'err' variable for the connection error
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fmt.Println("Database connection established")
}
