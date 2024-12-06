package donasi

import (
	"backend_relawanku/model"
	"time"

	"gorm.io/gorm"
)

func NewDonasiRepository(db *gorm.DB) *DonasiRepo {
	return &DonasiRepo{
		db: db,
	}
}

type DonasiRepo struct {
	db *gorm.DB
}

func (donasiRepo DonasiRepo) CreateDonasi(donasi model.Donasi) (model.Donasi, error) {
	donasiDb := FromModelDonasi(donasi)
	result := donasiRepo.db.Create(&donasiDb)
	if result.Error != nil {
		return model.Donasi{}, nil
	}
	return donasiDb.ToModelDonasi(), nil
}

func (donasiRepo DonasiRepo) UpdateDonasi(donasiId uint, donasi model.Donasi) (model.Donasi, error) {
	donasiDb := FromModelDonasi(donasi)
	result := donasiRepo.db.First(&donasiDb, donasiId)
	if result.Error != nil {
		return model.Donasi{}, result.Error
	}

	donasiDb.Title = donasi.Title
	donasiDb.Description = donasi.Description
	donasiDb.News = donasi.News
	donasiDb.TargetDonation = donasi.TargetDonation
	donasiDb.Category = donasi.Category
	donasiDb.ImageUrl = donasi.ImageUrl
	donasiDb.StartedAt = donasi.StartedAt
	donasiDb.FinishedAt = donasi.FinishedAt
	donasiDb.UpdatedAt = time.Now()

	saveResult := donasiRepo.db.Save(&donasiDb)
	if saveResult.Error != nil {
		return model.Donasi{}, saveResult.Error
	}

	return donasiDb.ToModelDonasi(), nil
}

func (donasiRepo DonasiRepo) DeleteDonasi(donasiId uint) error {
	var donasiDb Donasi
	result := donasiRepo.db.First(&donasiDb, donasiId)
	if result.Error != nil {
		return result.Error
	}

	deleteResult := donasiRepo.db.Delete(&donasiDb)
	return deleteResult.Error
}

func (donasiRepo DonasiRepo) GetAllDonasi() ([]model.Donasi, error) {
	var donasiDb []model.Donasi
	if err := donasiRepo.db.Find(&donasiDb).Error; err != nil {
		return nil, err
	}
	return donasiDb, nil
}

func (donasiRepo DonasiRepo) GetDonasiByCategory(category string) ([]model.Donasi, error) {
	var donasiDb []model.Donasi
	if err := donasiRepo.db.Where("category = ?", category).Find(&donasiDb).Error; err != nil {
		return nil, err
	}
	return donasiDb, nil
}

func (donasiRepo DonasiRepo) GetDonasiById(donasiId uint) (model.Donasi, error) {
	var donasiDb model.Donasi
	result := donasiRepo.db.First(&donasiDb, donasiId)
	if result.Error != nil {
		return model.Donasi{}, nil
	}
	return donasiDb, nil
}