package models

import (
	"gorm.io/gorm"
)

// Exercise model
type Exercise struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"` // in seconds
	Intensity   string    `json:"intensity"`
	Workouts    []Workout `gorm:"many2many:workout_exercises;" json:"workouts"`
}
