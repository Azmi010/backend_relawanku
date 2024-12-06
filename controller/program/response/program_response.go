package response

import "backend_relawanku/model"

type ProgramResponse struct {
	ID                uint   `json:"id" form:"id"`
	Title             string `json:"title" form:"title"`
	Quota    		  int    `json:"volunteer_quota" form:"quota"`
	Category          string `json:"category" form:"category"`
	Location          string `json:"location" form:"location"`
	ImageUrl          string  `json:"image_url" form:"image_url"`
	Details           string `json:"details" form:"details"`
}

func FromModel(program model.Program) ProgramResponse {
	return ProgramResponse{
		ID:             program.ID,
		Title:          program.Title,
		Quota: 			program.Quota,
		Category:       program.Category,
		Location:       program.Location,
		ImageUrl:       program.ImageUrl,
		Details:        program.Details,
	}
}
