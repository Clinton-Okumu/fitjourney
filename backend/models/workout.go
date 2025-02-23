package models

import (
	"gorm.io/gorm"
)

// Workout model
type Workout struct {
	gorm.Model
	Name      string     `json:"name"`
	UserID    uint       `json:"user_id"`
	Exercises []Exercise `gorm:"many2many:workout_exercises;" json:"exercises"`
}
