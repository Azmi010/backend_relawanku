package registration

import (
	"backend_relawanku/model"
	"strings"

	"gorm.io/gorm"
)

type UserProgramRepository struct {
	db *gorm.DB
}

func NewUserProgramRepository(db *gorm.DB) *UserProgramRepository {
	return &UserProgramRepository{db: db}
}

func (repo *UserProgramRepository) RegisterProgram(userProgram model.UserProgram) (model.UserProgram, error) {
	err := repo.db.Create(&userProgram).Error
	return userProgram, err
}

func (repo *UserProgramRepository) GetUserPrograms(userID uint) ([]model.Program, error) {
    var programs []model.Program
	err := repo.db.Joins("JOIN user_programs ON user_programs.program_id = programs.id").
		Where("user_programs.user_id = ?", userID).
		Find(&programs).Error
	return programs, err
}

func (repo *UserProgramRepository) FindUserIDByEmail(email string) (uint, error) {
    var user model.User
    err := repo.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return 0, err 
    }
    return user.ID, nil
}

func (repo *UserProgramRepository) FindProgramIDByName(name string) (uint, error) {
	var program model.Program
	err := repo.db.Where("LOWER(title) = ?", strings.ToLower(name)).First(&program).Error
	return program.ID, err
}
