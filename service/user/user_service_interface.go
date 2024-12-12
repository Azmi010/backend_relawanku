package user

import "backend_relawanku/model"

type UserServiceInterface interface {
	GetUserByID(userId uint) (model.User, error)
}