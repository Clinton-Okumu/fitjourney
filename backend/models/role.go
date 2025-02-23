package models

import (
	"gorm.io/gorm"
)

// Role model
type Role struct {
	gorm.Model
	Name string `gorm:"unique;size:50" json:"name"`
}
