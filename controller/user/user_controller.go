package user

import (
	"backend_relawanku/controller/base"
	response "backend_relawanku/controller/user/reponse"
	"backend_relawanku/controller/user/request"
	"backend_relawanku/service/user"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewUserController(us user.UserServiceInterface) *UserController {
	return &UserController{
		userServiceInterfae: us,
	}
}

type UserController struct {
	userServiceInterfae user.UserServiceInterface
}

// @Summary      Dapatkan User Sesuai ID
// @Description  Mengambil data user sesuai ID
// @Tags         profiles
// @Param 		 id path uint true "User ID"
// @Sec
// @Produce      json
// @Success      200  {array}   response.UserResponse
// @Router       /api/v1/user/profile/{id} [get]
// @Security     BearerAuth
func (userController UserController) GetUserByIDController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid donasi ID"))
	}

	donasi, err := userController.userServiceInterfae.GetUserByID(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.UserFromModel(donasi))
}

// @Summary      Update Profile
// @Description  Memperbarui profile berdasarkan ID
// @Tags         profiles
// @Accept       json
// @Produce      json
// @Param        id       path      uint                   true  "ID User"
// @Param        user  body      request.UserRequest  true  "Informasi Profile yang Diperbarui"
// @Success      200      {object}  map[string]interface{}
// @Router       /api/v1/user/profile/{id} [put]
// @Security     BearerAuth
func (userController UserController) UpdateUserController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid user ID"))
	}

	updateRequest := request.UserRequest{}
	if err := c.Bind(&updateRequest); err != nil {
		return base.ErrorResponse(c, err)
	}

	file, fileHeader, err := c.Request().FormFile("image_url")
	if err != nil {
		return base.ErrorResponse(c, errors.New("failed to read image file"))
	}
	defer file.Close()

	updatedUser, err := userController.userServiceInterfae.UpdateUser(uint(id), updateRequest.ToModelUser(), file, fileHeader)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.UserFromModel(updatedUser))
}

// @Summary      Update Profile
// @Description  Memperbarui profile berdasarkan ID
// @Tags         profiles
// @Accept       json
// @Produce      json
// @Param        id       path      uint                   true  "ID User"
// @Param        user  body      request.UpdatePasswordRequest  true  "Informasi Profile yang Diperbarui"
// @Success      200      {object}  map[string]interface{}
// @Router       /api/v1/user/profile/{id} [put]
// @Security     BearerAuth
func (userController UserController) UpdatePasswordController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid user ID"))
	}

	updatePasswordRequest := request.UpdatePasswordRequest{}
	if err := c.Bind(&updatePasswordRequest); err != nil {
		return base.ErrorResponse(c, err)
	}

	err = userController.userServiceInterfae.UpdatePassword(uint(id), updatePasswordRequest.OldPassword, updatePasswordRequest.NewPassword)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, "update password success")
}