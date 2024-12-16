package request

import "backend_relawanku/model"

type UserRequest struct {
	Username *string `json:"username" form:"username"`
	Gender   *string `json:"gender" form:"gender"`
	Address  *string `json:"address" form:"address"`
	ImageUrl *string `json:"image_url" form:"image_url"`
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
