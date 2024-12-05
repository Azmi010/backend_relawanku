package donasi

import (
	"backend_relawanku/model"
	"mime/multipart"
)

type DonasiServiceInterface interface {
	CreateDonasi(donasi model.Donasi, file multipart.File, fileHeader *multipart.FileHeader) (model.Donasi, error)
	UpdateDonasi(donasiId uint, donasi model.Donasi, file multipart.File, fileHeader *multipart.FileHeader) (model.Donasi, error)
	DeleteDonasi(donasiId uint) error
	GetAllDonasi() ([]model.Donasi, error)
}