package donasi

import (
	"backend_relawanku/model"
	"time"

	"gorm.io/gorm"
)

type Donasi struct {
	ID             uint                `gorm:"primaryKey"`
	Title          string              `gorm:"not null" json:"title" form:"title"`
	Description    string              `gorm:"not null" json:"description" form:"description"`
	Location       string              `gorm:"not null" json:"location" form:"location"`
	News           string              `gorm:"not null" json:"news" form:"news"`
	TargetDonation float64             `gorm:"not null" json:"target_donation" form:"target_donation"`
	Category       string              `gorm:"not null" json:"category" form:"category"`
	ImageUrl       string              `gorm:"not null" json:"image_url" form:"image_url"`
	StartedAt      string              `gorm:"not null" json:"started_at" form:"started_at"`
	FinishedAt     string              `gorm:"not null" json:"finished_at" form:"finished_at"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
	DeletedAt      gorm.DeletedAt      `gorm:"index" json:"deleted_at,omitempty"`
	Transactions   []model.Transaction `gorm:"foreignKey:DonasiID" json:"donasi_id"`
}

func FromModelDonasi(donasi model.Donasi) Donasi {
	return Donasi{
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
		CreatedAt:      donasi.CreatedAt,
		UpdatedAt:      donasi.UpdatedAt,
		DeletedAt:      donasi.DeletedAt,
	}
}

func (donasi Donasi) ToModelDonasi() model.Donasi {
	return model.Donasi{
		Model: gorm.Model{
			ID:        donasi.ID,
			CreatedAt: donasi.CreatedAt,
			UpdatedAt: donasi.UpdatedAt,
			DeletedAt: donasi.DeletedAt,
		},
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
