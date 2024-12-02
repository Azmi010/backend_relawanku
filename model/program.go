package model

import (
	"gorm.io/gorm"
	"time"
)

type Program struct {
	gorm.Model
	Title           string         `json:"title"`
	Quota           int            `json:"quota"`
	StartDate       time.Time      `json:"start_date"`
	EndDate         time.Time      `json:"end_date"`
	Category        string         `json:"category"`
	Location        string         `json:"location"`
	ImageUrl        string         `json:"image_url"`
	Details         string         `json:"details"`
}
