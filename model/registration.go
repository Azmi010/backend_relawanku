package model

import (
	"gorm.io/gorm"
)

type UserProgram struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	ProgramID   uint   `json:"program_id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	Motivation  string `json:"motivation"`
	PhoneNumber string `json:"phone_number"` 
}

