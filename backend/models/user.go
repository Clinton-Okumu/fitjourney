package models

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string    `gorm:"uniqueIndex;size:100" json:"username"`
	Password string    `json:"password"`
	Email    string    `gorm:"uniqueIndex;size:100" json:"email"`
	Workouts []Workout `gorm:"foreignKey:UserID" json:"workouts"`
}
