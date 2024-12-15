package transaction

import (
	"backend_relawanku/model"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint           `json:"id"`
	Nominal   float64        `json:"nominal"`
	Note      string         `json:"note"`
	DonasiID  uint           `json:"donasi_id"`
	UserID    uint           `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func FromModelTransaction(transaction model.Transaction) Transaction {
	return Transaction{
		ID:       transaction.ID,
		Nominal:  transaction.Nominal,
		Note:     transaction.Note,
		DonasiID: transaction.DonasiID,
		UserID:   transaction.UserID,
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.UpdatedAt,
		DeletedAt: transaction.DeletedAt,
	}
}

func (transaction Transaction) ToModelTransaction() model.Transaction {
	return model.Transaction{
		Model: gorm.Model{
			ID: transaction.ID,
			CreatedAt: transaction.CreatedAt,
			UpdatedAt: transaction.UpdatedAt,
			DeletedAt: transaction.DeletedAt,
		},
		Nominal: transaction.Nominal,
		Note: transaction.Note,
		DonasiID: transaction.DonasiID,
		UserID: transaction.UserID,
	}
}