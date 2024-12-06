package request

import (
	"backend_relawanku/model"
	"time"
)

type CreateProgramRequest struct {
	Title           string         `json:"title" form:"title"`
	Quota           int            `json:"quota" form:"quota"`
	StartDate       time.Time      `json:"start_date" form:"start_date"`
	EndDate         time.Time      `json:"end_date" form:"end_date"`
	Category        string         `json:"category" form:"category"`
	Location        string         `json:"location" form:"location"`
	ImageUrl        string         `json:"image_url" form:"image_url"`
	Details         string         `json:"details" form:"details"`
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
