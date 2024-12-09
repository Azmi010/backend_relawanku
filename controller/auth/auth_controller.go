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

// @Summary      Registrasi Pengguna
// @Description  Mendaftarkan pengguna baru
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        register  body      request.RegisterRequest  true  "Informasi Registrasi"
// @Success      201       {object}  map[string]interface{}
// @Router       /api/v1/register [post]
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

// @Summary      Login Pengguna
// @Description  Proses login pengguna
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login  body      request.LoginRequest  true  "Informasi Login"
// @Success      200    {object}  map[string]interface{}
// @Router       /api/v1/login [post]
func (authController AuthController) LoginController(c echo.Context) error {
	userLogin := request.LoginRequest{}
	c.Bind(&userLogin)
	user, token, err := authController.authServiceInterface.Login(userLogin.LoginToModelUser(), userLogin.LoginToModelAdmin())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.LoginFromModel(user, token))
}