package config

import (
	"backend/models"
	"fmt"
	"log"
)

// SeedRoles and Users function to populate the database with default roles and users
func SeedData() {
	// Seed Roles
	var role models.Role
	if err := DB.First(&role).Error; err != nil {
		// Create default roles (Admin and User)
		roles := []models.Role{
			{Name: "Admin"},
			{Name: "User"},
		}
		for _, r := range roles {
			if err := DB.Create(&r).Error; err != nil {
				log.Fatal("Error seeding roles:", err)
			}
		}
		fmt.Println("✅ Seeded default roles")
	}

	// Seed Users
	var user models.User
	if err := DB.First(&user).Error; err != nil {
		// Create default users with roles
		users := []models.User{
			{Username: "admin", Password: "adminpass", Email: "admin@example.com", RoleID: 1},     // Admin
			{Username: "john_doe", Password: "password123", Email: "john@example.com", RoleID: 2}, // User
		}
		for _, u := range users {
			if err := DB.Create(&u).Error; err != nil {
				log.Fatal("Error seeding users:", err)
			}
		}
		fmt.Println("✅ Seeded default users")
	}

	// Seed other models (Exercise and Workout)...
}
