package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"-"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	gorm.Model
	Name  string `gorm:"unique;not null" json:"name"`
	Users []User `gorm:"many2many:user_roles;"`
}
