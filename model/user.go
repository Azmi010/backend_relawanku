package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" form:"username"`
	Email    string `gorm:"email;not null" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
}