package auth

import "backend_relawanku/model"

type AuthRepository interface {
	Register(user model.User) (model.User, error)
	LoginUser(user model.User) (model.User, error)
	LoginAdmin(admin model.Admin) (model.Admin, error)
	IsUsernameOrEmailExists(username string, email string) (bool, error)
}