package models

import "gorm.io/gorm"

type Workout struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
}
