package program

import (
	"backend_relawanku/helper"
	"backend_relawanku/model"
	"backend_relawanku/repository/program"
	"errors"
	"mime/multipart"
	
)

type ProgramService struct {
	repo *program.ProgramRepository
}

func NewProgramService(repo *program.ProgramRepository) *ProgramService {
	return &ProgramService{repo: repo}
}

func (service *ProgramService) CreateProgram(program model.Program, file multipart.File, fileHeader *multipart.FileHeader) (model.Program, error) {
	imageURL, err := helper.UploadImageToFirebase("my-chatapp-01.appspot.com", "programs", fileHeader.Filename, file)
	if err != nil {
		return model.Program{}, errors.New(("failed to upload image to Firebase"))
	}

	program.ImageUrl = imageURL
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

func (service *ProgramService) UpdateProgram(id uint, updatedData model.Program, file multipart.File, fileHeader *multipart.FileHeader) (model.Program, error) {
	if file != nil && fileHeader != nil {
		imageURL, err := helper.UploadImageToFirebase("my-chatapp-01.appspot.com", "programs", fileHeader.Filename, file)
		if err != nil {
			return model.Program{}, errors.New("failed to upload image to Firebase")
		}
		updatedData.ImageUrl = imageURL
	}
	return service.repo.UpdateProgram(id, updatedData)
}

func (service *ProgramService) DeleteProgram(id uint) error {
	return service.repo.DeleteProgram(id)
}
