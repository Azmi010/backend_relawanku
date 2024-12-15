package registration

import (
	"backend_relawanku/model"
	"backend_relawanku/repository/registration"
	"errors"
)

type UserProgramService struct {
	repo registration.UserProgramRepositoryInterface
}

func NewUserProgramService(repo registration.UserProgramRepositoryInterface) *UserProgramService {
	return &UserProgramService{repo: repo}
}

func (service *UserProgramService) RegisterProgram(email, namaProgram, fullName, motivation, phoneNumber string) (model.UserProgram, error) {
	userID, err := service.repo.FindUserIDByEmail(email)
	if err != nil {
		return model.UserProgram{}, errors.New("user not found")
	}

	programID, err := service.repo.FindProgramIDByName(namaProgram)
	if err != nil {
		return model.UserProgram{}, errors.New("program not found")
	}

	userProgram := model.UserProgram{
		UserID:      userID,
		ProgramID:   programID,
		Email:       email,
		FullName:    fullName,
		Motivation:  motivation,
		PhoneNumber: phoneNumber, 
	}

	return service.repo.RegisterProgram(userProgram)
}

func (service *UserProgramService) GetUserPrograms(userID uint) ([]model.Program, error) {
    return service.repo.GetUserPrograms(userID)  // Pass userID ke repository
}

