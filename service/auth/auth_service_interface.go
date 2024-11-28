package auth

import "backend_relawanku/model"

type AuthServiceInterface interface {
	Register(user model.User) (model.User, error)
	Login(user model.User, admin model.Admin) (model.User, string, error)
}