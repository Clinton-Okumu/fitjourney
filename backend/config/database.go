package config

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
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Check if all required environment variables are set
	requiredEnvVars := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT"}
	for _, env := range requiredEnvVars {
		if os.Getenv(env) == "" {
			log.Fatalf("Please set the %s environment variable", env)
		}
	}

	// Prepare the Data Source Name (DSN) for PostgreSQL connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// Assign the connected database to the DB variable
	DB = db
	fmt.Println("✅ Database connected successfully")

	// Auto-migrate the User and Role models
	err = DB.AutoMigrate(&User{}, &Role{})
	if err != nil {
		log.Fatal("Error auto-migrating: ", err)
	}

	// Seed default roles
}
