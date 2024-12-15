package registration

import "backend_relawanku/model"

type UserProgramRepositoryInterface interface {
	RegisterProgram(userProgram model.UserProgram) (model.UserProgram, error)
	GetUserPrograms(userID uint) ([]model.Program, error)
	FindUserIDByEmail(email string) (uint, error)       
	FindProgramIDByName(name string) (uint, error)      
}
