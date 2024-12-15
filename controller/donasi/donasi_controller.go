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
        DonasiServiceInterface: ds,
    }
}

type DonasiController struct {
    DonasiServiceInterface donasi.DonasiServiceInterface
}

// @Summary      Buat Donasi Baru
// @Description  Membuat donasi baru oleh admin
// @Tags         donasi
// @Accept       json
// @Produce      json
// @Param        donasi  body      request.DonasiRequest  true  "Informasi Donasi"
// @Success      201      {object}  map[string]interface{}
// @Router       /api/v1/admin/donasi [post]
// @Security     BearerAuth
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

	createdDonasi, err := donasiController.DonasiServiceInterface.CreateDonasi(donasiCreated.DonasiToModel(), file, fileHeader)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.DonasiFromModel(createdDonasi))
}

// @Summary      Update Donasi
// @Description  Memperbarui donasi berdasarkan ID
// @Tags         donasi
// @Accept       json
// @Produce      json
// @Param        id       path      uint                   true  "ID Donasi"
// @Param        donasi  body      request.DonasiRequest  true  "Informasi Donasi yang Diperbarui"
// @Success      200      {object}  map[string]interface{}
// @Router       /api/v1/admin/donasi/{id} [put]
// @Security     BearerAuth
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

	updatedDonasi, err := donasiController.DonasiServiceInterface.UpdateDonasi(uint(id), donasiUpdated.DonasiToModel(), file, fileHeader)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.DonasiFromModel(updatedDonasi))
}

// @Summary      Hapus Donasi
// @Description  Menghapus donasi berdasarkan ID
// @Tags         donasi
// @Param        id  path      uint  true  "ID Donasi"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/admin/donasi/{id} [delete]
// @Security     BearerAuth
func (donasiController DonasiController) DeleteDonasiController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid donasi ID"))
	}

	err = donasiController.DonasiServiceInterface.DeleteDonasi(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, "Donasi deleted successfully")
}

// @Summary      Dapatkan Semua Donasi
// @Description  Mengambil daftar semua donasi
// @Tags         donasi
// @Produce      json
// @Success      200  {array}   map[string]interface{}
// @Router       /api/v1/admin/donasi [get]
// @Security     BearerAuth
func (donasiController DonasiController) GetAllDonasiController(c echo.Context) error {
	donasiList, err := donasiController.DonasiServiceInterface.GetAllDonasi()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, donasiList)
}

// @Summary      Dapatkan Donasi Sesuai Kategori
// @Description  Mengambil daftar semua donasi sesuai kategori
// @Tags         donasi
// @Param 		 category path string true "Category Donasi"
// @Produce      json
// @Success      200  {array}   map[string]interface{}
// @Router       /api/v1/admin/donasi/{category} [get]
// @Security     BearerAuth
func (donasiController DonasiController) GetDonasiByCategoryController(c echo.Context) error {
	category := c.Param("category")

	donasiList, err := donasiController.DonasiServiceInterface.GetDonasiByCategory(category)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, donasiList)
}

// @Summary      Dapatkan Donasi Sesuai ID
// @Description  Mengambil daftar semua donasi sesuai ID
// @Tags         donasi
// @Param 		 id path uint true "Donasi ID"
// @Sec
// @Produce      json
// @Success      200  {array}   map[string]interface{}
// @Router       /api/v1/admin/donasi/{id} [get]
// @Security     BearerAuth
func (donasiController DonasiController) GetDonasiByIdController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return base.ErrorResponse(c, errors.New("invalid donasi ID"))
	}

	donasi, err := donasiController.DonasiServiceInterface.GetDonasiById(uint(id))
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, donasi)
}