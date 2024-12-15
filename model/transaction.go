package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Nominal float64
	Note string
	DonasiID uint
	UserID uint
}