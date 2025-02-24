package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"unique;size:100"  json:"name"`
}
