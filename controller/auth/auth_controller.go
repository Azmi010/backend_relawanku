package auth

import (
	"backend_relawanku/controller/auth/request"
	"backend_relawanku/controller/auth/response"
	"backend_relawanku/controller/base"
	authService "backend_relawanku/service/auth"

	"github.com/labstack/echo/v4"
)

func NewAuthController(as authService.AuthServiceInterface) *AuthController {
	return &AuthController{
		authServiceInterface: as,
	}
}

type AuthController struct {
	authServiceInterface authService.AuthServiceInterface
}

func (authController AuthController) RegisterController(c echo.Context) error {
	userRegister := request.RegisterRequest{}
	if err := c.Bind(&userRegister); err != nil {
		return base.ErrorResponse(c, err)
	}

	user, err := authController.authServiceInterface.Register(userRegister.RegisterToModel())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.RegisterFromModel(user))
}
