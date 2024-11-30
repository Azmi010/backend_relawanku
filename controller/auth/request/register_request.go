package request

import "backend_relawanku/model"

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (registerRequest RegisterRequest) RegisterToModel() model.User {
	return model.User{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
}
