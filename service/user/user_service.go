package user

import (
	"backend_relawanku/model"
	"backend_relawanku/repository/user"
)

func NewUserService(ur user.UserRepository) *UserService {
	return &UserService{
		userRepoInterface: ur,
	}
}

type UserService struct {
	userRepoInterface user.UserRepository
}

func (userService UserService) GetUserByID(userId uint) (model.User, error) {
	detailUser, err := userService.userRepoInterface.GetUserByID(userId)
	if err != nil {
		return model.User{}, nil
	}
	return detailUser, nil
}