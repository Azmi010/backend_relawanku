package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title     string         `gorm:"not null" json:"title" form:"title"`
	SubTitle  string         `gorm:"not null" json:"sub_title" form:"sub_title"`
	Content   string         `gorm:"not null" json:"content" form:"content"`
	Category  string         `gorm:"not null" json:"category" form:"category"`
	View      int            `gorm:"not null" json:"view" form:"view"`
	ImageUrl  string         `gorm:"not null" json:"image_url" form:"image_url"`
}