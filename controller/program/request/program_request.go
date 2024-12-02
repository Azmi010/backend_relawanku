package request

import (
	"backend_relawanku/model"
	"time"
)

type CreateProgramRequest struct {
	Title           string         `json:"title"`
	Quota           int            `json:"quota"`
	StartDate       time.Time      `json:"start_date"`
	EndDate         time.Time      `json:"end_date"`
	Category        string         `json:"category"`
	Location        string         `json:"location"`
	ImageUrl        string         `json:"image_url"`
	Details         string         `json:"details"`
}

func (r CreateProgramRequest) ToModel() model.Program {
	return model.Program{
		Title:          r.Title,
		Quota:   		r.Quota,
		Category:       r.Category,
		Location:       r.Location,
		ImageUrl:       r.ImageUrl,
		Details:        r.Details,
	}
}
