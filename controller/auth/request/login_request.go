package request

import "backend_relawanku/model"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (loginRequest LoginRequest) LoginToModelUser() model.User {
	return model.User{
		Username:    loginRequest.Username,
		Password: loginRequest.Password,
	}
}

func (loginRequest LoginRequest) LoginToModelAdmin() model.Admin {
	return model.Admin{
		Username:    loginRequest.Username,
		Password: loginRequest.Password,
	}
}

// func (loginRequest LoginRequest) LoginToModel(userType string) interface{} {
// 	if userType == "admin" {
// 		return model.Admin{
// 			Username: loginRequest.Username,
// 			Password: loginRequest.Password,
// 		}
// 	}
// 	return model.User{
// 		Username: loginRequest.Username,
// 		Password: loginRequest.Password,
// 	}
// }