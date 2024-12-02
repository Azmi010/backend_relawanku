package program

import (
	"backend_relawanku/controller/base"
	"backend_relawanku/model"
	"backend_relawanku/service/program"

	"strconv"

	"github.com/labstack/echo/v4"
)

type ProgramController struct {
	service *program.ProgramService
}

func NewProgramController(service *program.ProgramService) *ProgramController {
	return &ProgramController{service: service}
}

func (ctrl *ProgramController) CreateProgram(c echo.Context) error {
	var program model.Program
	if err := c.Bind(&program); err != nil {
		return base.ErrorResponse(c, err)
	}
	createdProgram, err := ctrl.service.CreateProgram(program)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, createdProgram)
}

func (ctrl *ProgramController) GetAllPrograms(c echo.Context) error {
	programs, err := ctrl.service.GetAllPrograms()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, programs)
}

func (ctrl *ProgramController) GetProgramByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	program, err := ctrl.service.GetProgramByID(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, program)
}

func (ctrl *ProgramController) UpdateProgram(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	var updatedProgram model.Program
	if err := c.Bind(&updatedProgram); err != nil {
		return base.ErrorResponse(c, err)
	}

	program, err := ctrl.service.UpdateProgram(uint(id), updatedProgram)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, program)
}

func (ctrl *ProgramController) DeleteProgram(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	err = ctrl.service.DeleteProgram(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, "Program deleted successfully")
}

func (ctrl *ProgramController) GetProgramsByCategory(c echo.Context) error {
	category := c.Param("category")
	programs, err := ctrl.service.GetProgramsByCategory(category)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, programs)
}

func (ctrl *ProgramController) GetLatestProgram(c echo.Context) error {
	programs, err := ctrl.service.GetLatestProgram()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, programs)
}
