package transaction

import (
	"backend_relawanku/model"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID            uint           `json:"id"`
	TransactionID string         `gorm:"unique_index" json:"transaction_id"`
	Nominal       float64        `json:"nominal"`
	Note          string         `json:"note"`
	DonasiID      uint           `json:"donasi_id"`
	UserID        uint           `json:"user_id"`
	Status        string         `json:"status"`
	PaymentUrl    string         `json:"payment_url"`
	User          model.User     `json:"user"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func FromModelTransaction(transaction model.Transaction) Transaction {
	return Transaction{
		ID:            transaction.ID,
		TransactionID: transaction.TransactionID,
		Nominal:       transaction.Nominal,
		Note:          transaction.Note,
		DonasiID:      transaction.DonasiID,
		UserID:        transaction.UserID,
		Status:        transaction.Status,
		PaymentUrl:    transaction.PaymentUrl,
		User:          transaction.User,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
		DeletedAt:     transaction.DeletedAt,
	}
}

func (transaction Transaction) ToModelTransaction() model.Transaction {
	return model.Transaction{
		Model: gorm.Model{
			ID:        transaction.ID,
			CreatedAt: transaction.CreatedAt,
			UpdatedAt: transaction.UpdatedAt,
			DeletedAt: transaction.DeletedAt,
		},
		TransactionID: transaction.TransactionID,
		Nominal:       transaction.Nominal,
		Note:          transaction.Note,
		DonasiID:      transaction.DonasiID,
		UserID:        transaction.UserID,
		Status:        transaction.Status,
		PaymentUrl:    transaction.PaymentUrl,
		User:          transaction.User,
	}
}
