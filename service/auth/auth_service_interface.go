package auth

import "backend_relawanku/model"

type AuthServiceInterface interface {
	Register(user model.User) (model.User, error)
}