package donasi

import (
	"backend_relawanku/controller/base"
	"backend_relawanku/controller/donasi/request"
	"backend_relawanku/controller/donasi/response"
	"backend_relawanku/service/donasi"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewDonasiController(ds donasi.DonasiServiceInterface) *DonasiController {
	return &DonasiController{
		donasiServiceInterfae: ds,
	}
}

type DonasiController struct {
	donasiServiceInterfae donasi.DonasiServiceInterface
}

func (donasiController DonasiController) CreateDonasiController(c echo.Context) error {
	donasiCreated := request.DonasiRequest{}
	if err := c.Bind(&donasiCreated); err != nil {
		return base.ErrorResponse(c, err)
	}

	file, fileHeader, err := c.Request().FormFile("image_url")
	if err != nil {
		return base.ErrorResponse(c, errors.New("failed to read image file"))
	}
	defer file.Close()

	createdDonasi, err := donasiController.donasiServiceInterfae.CreateDonasi(donasiCreated.DonasiToModel(), file, fileHeader)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.DonasiFromModel(createdDonasi))
}

func (donasiController DonasiController) UpdateDonasiController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid donasi ID"))
	}

	donasiUpdated := request.DonasiRequest{}
	if err := c.Bind(&donasiUpdated); err != nil {
		return base.ErrorResponse(c, err)
	}

	file, fileHeader, err := c.Request().FormFile("image_url")
	if err != nil {
		return base.ErrorResponse(c, errors.New("failed to read image file"))
	}
	defer file.Close()

	updatedDonasi, err := donasiController.donasiServiceInterfae.UpdateDonasi(uint(id), donasiUpdated.DonasiToModel(), file, fileHeader)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.DonasiFromModel(updatedDonasi))
}

func (donasiController DonasiController) DeleteDonasiController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid donasi ID"))
	}

	err = donasiController.donasiServiceInterfae.DeleteDonasi(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, "Donasi deleted successfully")
}

func (donasiController DonasiController) GetAllDonasiController(c echo.Context) error {
	donasiList, err := donasiController.donasiServiceInterfae.GetAllDonasi()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, donasiList)
}