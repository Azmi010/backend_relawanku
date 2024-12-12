package user

import "backend_relawanku/model"

type UserRepository interface {
	GetUserByID(userId uint) (model.User, error)
}