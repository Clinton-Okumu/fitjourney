package models

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	Intensity   int       `json:"intensity"`
	Workout     []Workout `gorm:"many2many:workout_exercises;" json:"workout"`
}
