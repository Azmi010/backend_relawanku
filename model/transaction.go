package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TransactionID string
	Nominal       float64
	Note          string
	DonasiID      uint
	UserID        uint
	Status        string
	PaymentUrl    string
	User          User
}
