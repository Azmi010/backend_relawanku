package model

import "gorm.io/gorm"

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

type User struct {
	gorm.Model
	Username  string
	Email     string
	Password  string
	Role      string
	Gender    string
	Address   string
	ImageUrl string
}
