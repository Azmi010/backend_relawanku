package model

import "gorm.io/gorm"

type Donasi struct {
	gorm.Model
	Title          string
	Description    string
	Location       string
	News           string
	TargetDonation float64
	Category string
	ImageUrl string
	StartedAt string
	FinishedAt string
}
