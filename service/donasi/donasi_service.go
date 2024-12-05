package donasi

import (
	"backend_relawanku/helper"
	"backend_relawanku/model"
	donasiRepo "backend_relawanku/repository/donasi"
	"errors"
	"mime/multipart"
)

func NewDonasiService(dr donasiRepo.DonasiRepository) *DonasiService {
	return &DonasiService{
		donasiRepoInterface: dr,
	}
}

type DonasiService struct {
	donasiRepoInterface donasiRepo.DonasiRepository
}

func (donasiService DonasiService) CreateDonasi(donasi model.Donasi, file multipart.File, fileHeader *multipart.FileHeader) (model.Donasi, error) {
	imageURL, err := helper.UploadImageToFirebase("my-chatapp-01.appspot.com", "donasi", fileHeader.Filename, file)
	if err != nil {
		return model.Donasi{}, errors.New("failed to upload image to Firebase")
	}

	donasi.ImageUrl = imageURL

	createdDonasi, err := donasiService.donasiRepoInterface.CreateDonasi(donasi)
	if err != nil {
		return model.Donasi{}, errors.New("failed to create donasi")
	}

	return createdDonasi, nil
}

func (donasiService DonasiService) UpdateDonasi(donasiId uint, donasi model.Donasi, file multipart.File, fileHeader *multipart.FileHeader) (model.Donasi, error) {
	if file != nil && fileHeader != nil {
		imageURL, err := helper.UploadImageToFirebase("my-chatapp-01.appspot.com", "donasi", fileHeader.Filename, file)
		if err != nil {
			return model.Donasi{}, errors.New("failed to upload image to Firebase")
		}
		donasi.ImageUrl = imageURL
	}

	updated, err := donasiService.donasiRepoInterface.UpdateDonasi(donasiId, donasi)
	if err != nil {
		return model.Donasi{}, errors.New("failed to update donasi")
	}
	return updated, nil
}

func (donasiService DonasiService) DeleteDonasi(donasiId uint) error {
	err := donasiService.donasiRepoInterface.DeleteDonasi(donasiId)
	if err != nil {
		return errors.New("failed to delete donasi")
	}
	return nil
}

func (donasiService DonasiService) GetAllDonasi() ([]model.Donasi, error) {
	listDonasi, err := donasiService.donasiRepoInterface.GetAllDonasi()
	if err != nil {
		return []model.Donasi{}, nil
	}
	return listDonasi, nil
}
