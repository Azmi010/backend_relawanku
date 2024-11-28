package auth

import (
	"backend_relawanku/model"

	"gorm.io/gorm"
)

func NewAuthRepository(db *gorm.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

type AuthRepo struct {
	db *gorm.DB
}

func (authRepo AuthRepo) Register(user model.User) (model.User, error) {
	userDb := FromModel(user)
	result := authRepo.db.Create(&userDb)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return userDb.ToModel(), nil
}