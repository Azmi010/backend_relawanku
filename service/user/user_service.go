package user

import (
	"backend_relawanku/helper"
	"backend_relawanku/model"
	"backend_relawanku/repository/user"
	"errors"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
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

func (userService UserService) UpdateUser(userId uint, user model.User, file multipart.File, fileHeader *multipart.FileHeader) (model.User, error) {
	updated, err := userService.userRepoInterface.GetUserByID(userId)
    if err != nil {
        return model.User{}, errors.New("failed to find user")
    }

    if user.Username != "" {
        updated.Username = user.Username
    }
    if user.Gender != "" {
        updated.Gender = user.Gender
    }
    if user.Address != "" {
        updated.Address = user.Address
    }
	if file != nil && fileHeader != nil {
		imageURL, err := helper.UploadImageToFirebase("my-chatapp-01.appspot.com", "profiles", fileHeader.Filename, file)
		if err != nil {
			return model.User{}, errors.New("failed to upload image to Firebase")
		}
		updated.ImageUrl = imageURL
	}

    saved, err := userService.userRepoInterface.UpdateUser(userId, updated)
    if err != nil {
        return model.User{}, errors.New("failed to update user")
    }

    return saved, nil
}

func (userService *UserService) UpdatePassword(userId uint, oldPassword string, newPassword string) error {
	user, err := userService.userRepoInterface.GetUserByID(userId)
	if err != nil {
		return errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("old password does not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	err = userService.userRepoInterface.UpdatePassword(userId, string(hashedPassword))
	if err != nil {
		return errors.New("failed to update password")
	}

	return nil
}