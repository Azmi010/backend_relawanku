package response

import "backend_relawanku/model"

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func RegisterFromModel(user model.User) RegisterResponse {
	return RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
