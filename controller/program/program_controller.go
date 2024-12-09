package program

import (
	"backend_relawanku/controller/base"
	"backend_relawanku/model"
	"backend_relawanku/service/program"
	"errors"
	"time"

	"strconv"

	"github.com/labstack/echo/v4"
)

type ProgramController struct {
	Service *program.ProgramService
}

func NewProgramController(service *program.ProgramService) *ProgramController {
	return &ProgramController{Service: service}
}

// @Summary      Buat Program Baru
// @Description  Membuat program baru oleh admin
// @Tags         programs
// @Accept       json
// @Produce      json
// @Param        program  body      request.CreateProgramRequest  true  "Informasi Program"
// @Success      201      {object}  map[string]interface{}
// @Router       /api/v1/admin/program [post]
// @Security     BearerAuth
func (ctrl *ProgramController) CreateProgram(c echo.Context) error {
	var program model.Program
	program.Title = c.FormValue("title")
	program.Category = c.FormValue("category")
	program.Location = c.FormValue("location")
	program.Details = c.FormValue("details")

	quota, _ := strconv.Atoi(c.FormValue("quota"))
	program.Quota = quota

	startDate, _ := time.Parse("2006-01-02", c.FormValue("start_date"))
	endDate, _ := time.Parse("2006-01-02", c.FormValue("end_date"))
	program.StartDate = startDate
	program.EndDate = endDate

	file, fileHeader, err := c.Request().FormFile("image_url")
	if err != nil {
		return base.ErrorResponse(c, errors.New("failed to read image file"))
	}
	defer file.Close()

	createdProgram, err := ctrl.Service.CreateProgram(program, file, fileHeader)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, createdProgram)
}

// @Summary      Dapatkan Semua Program
// @Description  Mengambil daftar semua program
// @Tags         programs
// @Produce      json
// @Success      200  {array}   response.ProgramResponse
// @Router       /api/v1/admin/programs [get]
// @Security     BearerAuth
func (ctrl *ProgramController) GetAllPrograms(c echo.Context) error {
	programs, err := ctrl.Service.GetAllPrograms()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, programs)
}

// @Summary      Dapatkan Artikel Sesuai ID
// @Description  Mengambil daftar semua artikel sesuai ID
// @Tags         programs
// @Param 		 id path uint true "Program ID"
// @Sec
// @Produce      json
// @Success      200  {array}   response.ProgramResponse
// @Router       /api/v1/admin/program/{id} [get]
// @Security     BearerAuth
func (ctrl *ProgramController) GetProgramByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	program, err := ctrl.Service.GetProgramByID(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, program)
}

// @Summary      Update Program
// @Description  Memperbarui program berdasarkan ID
// @Tags         programs
// @Accept       json
// @Produce      json
// @Param        id       path      uint                   true  "ID Program"
// @Param        program  body      model.Program  true  "Informasi Program yang Diperbarui"
// @Success      200      {object}  map[string]interface{}
// @Router       /api/v1/admin/program/{id} [put]
// @Security     BearerAuth
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

	file, fileHeader, err := c.Request().FormFile("image_url")
	if err != nil {
		return base.ErrorResponse(c, errors.New("failed to read image file"))
	}
	defer file.Close()

	program, err := ctrl.Service.UpdateProgram(uint(id), updatedProgram, file, fileHeader)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, program)
}

// @Summary      Hapus Program
// @Description  Menghapus program berdasarkan ID
// @Tags         programs
// @Param        id  path      uint  true  "ID Program"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/admin/program/{id} [delete]
// @Security     BearerAuth
func (ctrl *ProgramController) DeleteProgram(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	err = ctrl.Service.DeleteProgram(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, "Program deleted successfully")
}

// @Summary      Dapatkan Program Sesuai Kategori
// @Description  Mengambil daftar semua program sesuai kategori
// @Tags         programs
// @Param 		 category path string true "Category Program"
// @Produce      json
// @Success      200  {array}   response.ProgramResponse
// @Router       /api/v1/admin/program/{category} [get]
// @Security     BearerAuth
func (ctrl *ProgramController) GetProgramsByCategory(c echo.Context) error {
	category := c.Param("category")
	programs, err := ctrl.Service.GetProgramsByCategory(category)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, programs)
}

// @Summary      Dapatkan Program Terbaru
// @Description  Mengambil daftar semua program kategori
// @Tags         programs
// @Param 		 category path string true "Category Program"
// @Produce      json
// @Success      200  {array}   response.ProgramResponse
// @Router       /api/v1/admin/program/{latest} [get]
// @Security     BearerAuth
func (ctrl *ProgramController) GetLatestProgram(c echo.Context) error {
	programs, err := ctrl.Service.GetLatestProgram()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, programs)
}
