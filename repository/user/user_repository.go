package user

import (
	"backend_relawanku/model"
	"backend_relawanku/repository/auth"

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

func (userRepo UserRepo) UpdateUser(userId uint, user model.User) (model.User, error) {
	userDb := auth.FromModelUser(user)
	result := userRepo.db.First(&userDb, userId)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	userDb.Username = user.Username
	userDb.Gender = user.Gender
	userDb.Address = user.Address

	saveResult := userRepo.db.Save(&userDb)
	if saveResult.Error != nil {
		return model.User{}, saveResult.Error
	}

	return userDb.ToModelUser(), nil
}

func (userRepo UserRepo) UpdatePassword(userId uint, newPassword string) error {
	return userRepo.db.Model(&model.User{}).Where("id = ?", userId).Update("password", newPassword).Error
}

// GetAllUsers mengambil semua user
func (userRepo UserRepo) GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := userRepo.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// DeleteUser menghapus user berdasarkan ID
func (userRepo UserRepo) DeleteUser(userId uint) error {
	result := userRepo.db.Delete(&model.User{}, userId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
