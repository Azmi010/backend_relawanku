package donasi

import "backend_relawanku/model"

type DonasiRepository interface {
	CreateDonasi(donasi model.Donasi) (model.Donasi, error)
	UpdateDonasi(donasiId uint, donasi model.Donasi) (model.Donasi, error)
	DeleteDonasi(donasiId uint) error
	GetAllDonasi() ([]model.Donasi, error)
}