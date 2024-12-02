package program

import (
	"backend_relawanku/model"
	"backend_relawanku/repository/program"
)

type ProgramService struct {
	repo *program.ProgramRepository
}

func NewProgramService(repo *program.ProgramRepository) *ProgramService {
	return &ProgramService{repo: repo}
}

func (service *ProgramService) CreateProgram(program model.Program) (model.Program, error) {
	return service.repo.CreateProgram(program)
}

func (service *ProgramService) GetAllPrograms() ([]model.Program, error) {
	return service.repo.GetAllPrograms()
}

func (service *ProgramService) GetProgramByID(id uint) (model.Program, error) {
	return service.repo.GetProgramByID(id)
}

func (service *ProgramService) GetProgramsByCategory(category string) ([]model.Program, error) {
	return service.repo.GetProgramsByCategory(category)
}

func (service *ProgramService) GetLatestProgram() (model.Program, error) {
	return service.repo.GetLatestProgram()
}

func (service *ProgramService) UpdateProgram(id uint, updatedData model.Program) (model.Program, error) {
	return service.repo.UpdateProgram(id, updatedData)
}

func (service *ProgramService) DeleteProgram(id uint) error {
	return service.repo.DeleteProgram(id)
}
