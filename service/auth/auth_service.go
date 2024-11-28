package auth

import (
	"backend_relawanku/model"
	authRepo "backend_relawanku/repository/auth"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(ar authRepo.AuthRepository) *AuthService {
	return &AuthService{
		authRepoInterface: ar,
	}
}

type AuthService struct {
	authRepoInterface authRepo.AuthRepository
}

func (authService AuthService) Register(user model.User) (model.User, error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return model.User{}, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	createdUser, err := authService.authRepoInterface.Register(user)
	if err != nil {
		return model.User{}, errors.New("failed to create user")
	}

	return createdUser, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}