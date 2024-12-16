package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title     string
	SubTitle  string
	Content   string
	Category  string
	View      int
	ImageUrl  string
}