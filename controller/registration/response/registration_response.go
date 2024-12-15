package response

import "backend_relawanku/model"

type RegisterProgramResponse struct {
	UserID      uint   `json:"user_id"`
	ProgramID   uint   `json:"program_id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	Motivation  string `json:"motivation"`
	PhoneNumber string `json:"phone_number"` 
}

func FromModel(userProgramModel interface{}) RegisterProgramResponse {
	model, ok := userProgramModel.(model.UserProgram)
	if !ok {
		panic("Invalid model type")
	}

	return RegisterProgramResponse{
		UserID:      model.UserID,
		ProgramID:   model.ProgramID,
		Email:       model.Email,
		FullName:    model.FullName,
		Motivation:  model.Motivation,
		PhoneNumber: model.PhoneNumber, 
	}
}

