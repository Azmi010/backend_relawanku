package registration

import (
	"backend_relawanku/controller/base"
	"backend_relawanku/controller/registration/request"
	"backend_relawanku/controller/registration/response"
	"backend_relawanku/service/registration"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserProgramController struct {
	service *registration.UserProgramService
}

func NewUserProgramController(service *registration.UserProgramService) *UserProgramController {
	return &UserProgramController{service: service}
}

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

func (ctrl *UserProgramController) GetUserPrograms(c echo.Context) error {
    // Mengambil userID dari context
    userID, ok := c.Get("userID").(uint)
    if !ok {
        return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
    }

    programs, err := ctrl.service.GetUserPrograms(userID)
    if err != nil {
        return base.ErrorResponse(c, err)
    }

    return base.SuccessResponse(c, programs)
}



