package program

import (
	"backend_relawanku/model"

	"gorm.io/gorm"
)

type ProgramRepository struct {
	db *gorm.DB
}

func NewProgramRepository(db *gorm.DB) *ProgramRepository {
	return &ProgramRepository{db: db}
}

func (repo *ProgramRepository) CreateProgram(program model.Program) (model.Program, error) {
	result := repo.db.Create(&program)
	return program, result.Error
}

func (repo *ProgramRepository) GetAllPrograms() ([]model.Program, error) {
	var programs []model.Program
	result := repo.db.Find(&programs)
	return programs, result.Error
}

func (repo *ProgramRepository) GetProgramByID(id uint) (model.Program, error) {
	var program model.Program
	result := repo.db.First(&program, id)
	return program, result.Error
}

func (repo *ProgramRepository) GetProgramsByCategory(category string) ([]model.Program, error) {
	var programs []model.Program
	result := repo.db.Where("category = ?", category).Find(&programs)
	return programs, result.Error
}

func (repo *ProgramRepository) GetLatestProgram() (model.Program, error) {
	var program model.Program
	result := repo.db.Order("created_at desc").First(&program)
	return program, result.Error
}

func (repo *ProgramRepository) UpdateProgram(id uint, updatedData model.Program) (model.Program, error) {
	var program model.Program
	err := repo.db.First(&program, id).Error
	if err != nil {
		return program, err
	}
	err = repo.db.Model(&program).Updates(updatedData).Error
	return program, err
}

func (repo *ProgramRepository) DeleteProgram(id uint) error {
	return repo.db.Delete(&model.Program{}, id).Error
}
