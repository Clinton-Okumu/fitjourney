package main

import (
	"backend/config"
	"backend/models"
	"fmt"
	"log"
)

func RunMigrations() {
	// Step 1: Connect to the database using your existing config function
	config.ConnectDB()
	db := config.DB // Get the DB connection

	// Step 2: Run migrations to create the tables (User and Role)
	err := db.AutoMigrate(&models.User{}, &models.Role{})
	if err != nil {
		log.Fatal("Migration failed:", err) // Handle any errors during migration
	}

	// Step 3: Confirm successful migration
	fmt.Println("Database migrated successfully!")
}
