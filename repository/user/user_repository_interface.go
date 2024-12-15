package user

import "backend_relawanku/model"

type UserRepository interface {
	GetUserByID(userId uint) (model.User, error)
	UpdateUser(userId uint, user model.User) (model.User, error)
	UpdatePassword(userId uint, newPassword string) error
	GetAllUsers() ([]model.User, error)           
	DeleteUser(userId uint) error  
}