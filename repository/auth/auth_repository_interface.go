package auth

import "backend_relawanku/model"

type AuthRepository interface {
	Register(user model.User) (model.User, error)
}