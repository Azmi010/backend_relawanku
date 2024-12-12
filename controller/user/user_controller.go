package user

import (
	"backend_relawanku/controller/base"
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
	return base.SuccessResponse(c, donasi)
}