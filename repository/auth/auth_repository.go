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
	userDb := FromModelUser(user)
	result := authRepo.db.Create(&userDb)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return userDb.ToModelUser(), nil
}

func (authRepo AuthRepo) LoginUser(user model.User) (model.User, error) {
	userDb := FromModelUser(user)
	result := authRepo.db.First(&userDb, "username = ?", user.Username)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return userDb.ToModelUser(), nil
}

func (authRepo AuthRepo) LoginAdmin(admin model.Admin) (model.Admin, error) {
	userDb := FromModelAdmin(admin)
	result := authRepo.db.First(&userDb, "username = ?", admin.Username)
	if result.Error != nil {
		return model.Admin{}, result.Error
	}
	return userDb.ToModelAdmin(), nil
}

func (authRepo AuthRepo) IsUsernameOrEmailExists(username string, email string) (bool, error) {
	var userCount int64
	var adminCount int64

	if err := authRepo.db.Model(&model.User{}).Where("username = ? OR email = ?", username, email).Count(&userCount).Error; err != nil {
		return false, err
	}

	if err := authRepo.db.Model(&model.Admin{}).Where("username = ? OR email = ?", username, email).Count(&adminCount).Error; err != nil {
		return false, err
	}

	if userCount > 0 || adminCount > 0 {
		return true, nil
	}

	return false, nil
}