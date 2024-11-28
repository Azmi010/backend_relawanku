package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string
	Email    string
	Password string
}