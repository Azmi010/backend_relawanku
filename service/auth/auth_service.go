package auth

import (
	"backend_relawanku/middleware"
	"backend_relawanku/model"
	authRepo "backend_relawanku/repository/auth"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(ar authRepo.AuthRepository, jt middleware.JwtInterface) *AuthService {
	return &AuthService{
		authRepoInterface: ar,
		jwtInterface:      jt,
	}
}

type AuthService struct {
	authRepoInterface authRepo.AuthRepository
	jwtInterface      middleware.JwtInterface
}

func (authService AuthService) Register(user model.User) (model.User, error) {
	isExists, err := authService.authRepoInterface.IsUsernameOrEmailExists(user.Username, user.Email)
	if err != nil {
		return model.User{}, err
	}

	if isExists {
		return model.User{}, errors.New("username or email already exists")
	}
	
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

func (authService AuthService) Login(user model.User, admin model.Admin) (model.User, string, error) {
	storedUser, err := authService.authRepoInterface.LoginUser(user)
	if err != nil {
		storedAdmin, err := authService.authRepoInterface.LoginAdmin(admin)
		if err != nil {
			return model.User{}, "", errors.New("admin or user not found")
		}

		if !CheckPasswordHash(admin.Password, storedAdmin.Password) {
			return model.User{}, "", errors.New("invalid credentials")
		}

		mappedAdmin := model.User{
			Model:    storedAdmin.Model,
			Username: storedAdmin.Username,
			Email:    storedAdmin.Email,
			Role:     string(model.RoleAdmin),
		}

		token, err := authService.jwtInterface.GenerateJWT(storedAdmin.Username, model.RoleAdmin)
		if err != nil {
			return model.User{}, "", err
		}

		return mappedAdmin, token, nil
	}

	if !CheckPasswordHash(user.Password, storedUser.Password) {
		return model.User{}, "", errors.New("invalid credentials")
	}

	token, err := authService.jwtInterface.GenerateJWT(storedUser.Username, model.UserRole(storedUser.Role))
	if err != nil {
		return model.User{}, "", err
	}

	return storedUser, token, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}