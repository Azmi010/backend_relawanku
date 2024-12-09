package program

import "backend_relawanku/model"

type ProgramRepositoryInterface interface {
	CreateProgram(program model.Program) (model.Program, error)
	GetAllPrograms() ([]model.Program, error)
	GetProgramByID(id uint) (model.Program, error)
	GetProgramsByCategory(category string) ([]model.Program, error)
	GetLatestProgram() (model.Program, error)
	UpdateProgram(id uint, updatedData model.Program) (model.Program, error)
	DeleteProgram(id uint) error
}
