package models

import (
	"fmt"
)

// SeedRoles inserts default roles into the database
func SeedRoles() {
	roles := []Role{
		{Name: "Admin"},
		{Name: "Workouter"},
		{Name: "User"},
	}

	for _, role := range roles {
		var existing Role
		result := DB.Where("name = ?", role.Name).First(&existing)
		if result.Error != nil {
			// Role not found, create it
			DB.Create(&role)
			fmt.Println("✅ Role added:", role.Name)
		} else {
			fmt.Println("⚠️ Role already exists:", role.Name)
		}
	}
}
