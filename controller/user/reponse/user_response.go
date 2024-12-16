package response

import (
	"backend_relawanku/model"
	"time"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClientsResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	RegTime   time.Time `json:"tanggal_registrasi"` 
}

func UserFromModel(user model.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Gender:    user.Gender,
		Address:   user.Address,
		ImageUrl:  user.ImageUrl,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func Clients(user model.User) ClientsResponse {
	return ClientsResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		RegTime:  user.CreatedAt,
	}
}
