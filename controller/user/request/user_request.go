package request

import "backend_relawanku/model"

type UserRequest struct {
	Username *string `json:"username" form:"username"`
	Gender   *string `json:"gender" form:"username"`
	Address  *string `json:"address" form:"username"`
	ImageUrl *string `json:"image_url" form:"username"`
}

func (userRequest *UserRequest) ToModelUser() model.User {
	return model.User{
		Username: ifNotNil(userRequest.Username, ""),
		Gender:   ifNotNil(userRequest.Gender, ""),
		Address:  ifNotNil(userRequest.Address, ""),
		ImageUrl: ifNotNil(userRequest.ImageUrl, ""),
	}
}

func ifNotNil(value *string, fallback string) string {
	if value != nil {
		return *value
	}
	return fallback
}
