package registration

import (
	"backend_relawanku/controller/base"
	"backend_relawanku/controller/registration/request"
	"backend_relawanku/controller/registration/response"
	"backend_relawanku/service/registration"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserProgramController struct {
	service *registration.UserProgramService
}

func NewUserProgramController(service *registration.UserProgramService) *UserProgramController {
	return &UserProgramController{service: service}
}

// @Summary      Daftar Program
// @Description  Mendaftar Pada Sebuah Program
// @Tags         programs
// @Accept       json
// @Produce      json
// @Param        program  body      request.RegisterProgramRequest  true  "Daftar Program"
// @Success      201      {object}  map[string]interface{}
// @Router       /api/v1/user/register-program [post]
// @Security     BearerAuth
func (ctrl *UserProgramController) RegisterProgram(c echo.Context) error {
	var req request.RegisterProgramRequest
	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	userProgram, err := ctrl.service.RegisterProgram(req.Email, req.NamaProgram, req.FullName, req.Motivation, req.PhoneNumber) // Tambahkan phoneNumber
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	res := response.FromModel(userProgram)
	return base.SuccessResponse(c, res)
}

// @Summary      Dapatkan Program Sesuai User ID
// @Description  Mengambil data program yang diikuti sesuai User ID
// @Tags         programs
// @Param 		 id path uint true "User ID"
// @Sec
// @Produce      json
// @Success      200  {array}   map[string]interface{}
// @Router       /api/v1/user/my-program/{id} [get]
// @Security     BearerAuth
func (ctrl *UserProgramController) GetUserPrograms(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	programs, err := ctrl.service.GetUserPrograms(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, programs)
}