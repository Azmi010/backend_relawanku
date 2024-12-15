package user

import (
	"backend_relawanku/model"
	"mime/multipart"
)

type UserServiceInterface interface {
	GetUserByID(userId uint) (model.User, error)
	UpdateUser(userId uint, user model.User, file multipart.File, fileHeader *multipart.FileHeader) (model.User, error)
	UpdatePassword(userId uint, oldPassword string, newPassword string) error
}