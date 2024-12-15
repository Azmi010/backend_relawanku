package request

type RegisterProgramRequest struct {
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	NamaProgram string `json:"nama_program"`
	Motivation  string `json:"motivation"`
	PhoneNumber string `json:"phone_number"` 
}

