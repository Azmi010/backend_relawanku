package response

import (
	"backend_relawanku/model"
)

type DonasiResponse struct {
	ID             uint    `json:"id" form:"id"`
	Title          string  `json:"title" form:"title"`
	Description    string  `json:"description" form:"description"`
	Location       string  `json:"location" form:"location"`
	News           string  `json:"news" form:"news"`
	TargetDonation float64 `json:"target_donation" form:"target_donation"`
	Category       string  `json:"category" form:"category"`
	ImageUrl       string  `json:"image_url" form:"image_url"`
	StartedAt      string  `json:"started_at" form:"started_at"`
	FinishedAt     string  `json:"finished_at" form:"finished_at"`
}

func DonasiFromModel(donasi model.Donasi) DonasiResponse {
	return DonasiResponse{
		ID:             donasi.ID,
		Title:          donasi.Title,
		Description:    donasi.Description,
		Location:       donasi.Location,
		News:           donasi.News,
		TargetDonation: donasi.TargetDonation,
		Category:       donasi.Category,
		ImageUrl:       donasi.ImageUrl,
		StartedAt:      donasi.StartedAt,
		FinishedAt:     donasi.FinishedAt,
	}
}
