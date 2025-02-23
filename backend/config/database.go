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
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using default environment")
	}

	// Get database credentials from .env
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error connecting to database:", err)
	}

	DB = db
	fmt.Println("✅ Connected to database")

	// Run migrations
	err = DB.AutoMigrate(&User{}, &Role{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	fmt.Println("✅ Database migrated successfully!")

	// Seed default roles
	SeedRoles()
}
