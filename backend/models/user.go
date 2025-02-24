package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;size:100"  json:"username"`
	Email    string `gorm:"unique;size:100"  json:"email"`
	Password string `json:"-"`
	RoleID   uint   `json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID" json:"role"`
}
