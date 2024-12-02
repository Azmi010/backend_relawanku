package response

import "backend_relawanku/model"

type ProgramResponse struct {
	ID                uint   `json:"id"`
	Title             string `json:"title"`
	Quota    		  int    `json:"volunteer_quota"`
	Category          string `json:"category"`
	Location          string `json:"location"`
	ImageUrl          string  `json:"image_url"`
	Details           string `json:"details"`
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
