package transaction

import (
	"backend_relawanku/model"
	"errors"

	"gorm.io/gorm"
)

func NewTransactionRepository(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{
		db: db,
	}
}

type TransactionRepo struct {
	db *gorm.DB
}

func (transactionRepo *TransactionRepo) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	if transaction.UserID == 0 || transaction.DonasiID == 0 {
		return model.Transaction{}, errors.New("user_id or donasi_id is missing")
	}
	if err := transactionRepo.db.Create(&transaction).Error; err != nil {
		return model.Transaction{}, err
	}

	var createdTransaction model.Transaction
	result := transactionRepo.db.Preload("User").First(&createdTransaction, transaction.ID)
	if result.Error != nil {
		return model.Transaction{}, result.Error
	}
	return createdTransaction, nil
}