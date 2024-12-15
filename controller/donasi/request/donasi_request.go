package request

import (
	"backend_relawanku/model"
)

type DonasiRequest struct {
	Title          string  `json:"title" form:"title"`
	Description    string  `json:"description" form:"description"`
	News           string  `json:"news" form:"news"`
	TargetDonation float64 `json:"target_donation" form:"target_donation"`
	Category       string  `json:"category" form:"category"`
	ImageUrl       string  `json:"image_url" form:"image_url"`
	StartedAt      string  `json:"started_at" form:"started_at"`
	FinishedAt     string  `json:"finished_at" form:"finished_at"`
}

func (donasiRequest DonasiRequest) DonasiToModel() model.Donasi {
	return model.Donasi{
		Title:          donasiRequest.Title,
		Description:    donasiRequest.Description,
		News:           donasiRequest.News,
		TargetDonation: donasiRequest.TargetDonation,
		Category:       donasiRequest.Category,
		ImageUrl:       donasiRequest.ImageUrl,
		StartedAt:      donasiRequest.StartedAt,
		FinishedAt:     donasiRequest.FinishedAt,
	}
}
