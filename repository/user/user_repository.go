package user

import (
	"backend_relawanku/model"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

type UserRepo struct {
	db *gorm.DB
}

func (userRepo UserRepo) GetUserByID(userId uint) (model.User, error) {
	var userDb model.User
	result := userRepo.db.First(&userDb, userId)
	if result.Error != nil {
		return model.User{}, nil
	}
	return userDb, nil
}