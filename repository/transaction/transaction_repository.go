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

func (transactionRepo *TransactionRepo) GetTransactionByID(transactionID int) (model.Transaction, error) {
	var transaction model.Transaction
	if err := transactionRepo.db.Preload("User").First(&transaction, transactionID).Error; err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}

func (transactionRepo *TransactionRepo) GetAllTransactions() ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := transactionRepo.db.Preload("User").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (transactionRepo *TransactionRepo) UpdateTransaction(transactionID int, updates map[string]interface{}) (model.Transaction, error) {
	if err := transactionRepo.db.Model(&model.Transaction{}).Where("id = ?", transactionID).Updates(updates).Error; err != nil {
		return model.Transaction{}, err
	}

	var updatedTransaction model.Transaction
	if err := transactionRepo.db.Preload("User").First(&updatedTransaction, transactionID).Error; err != nil {
		return model.Transaction{}, err
	}
	return updatedTransaction, nil
}

func (transactionRepo *TransactionRepo) UpdateTransactionStatus(transactionID int, status string) error {
	return transactionRepo.db.Model(&model.Transaction{}).Where("id = ?", transactionID).Update("status", status).Error
}

func (transactionRepo *TransactionRepo) DeleteTransaction(transactionID int) error {
	result := transactionRepo.db.Where("id = ?", transactionID).Delete(&model.Transaction{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("transaction not found")
	}
	return nil
}

func (transactionRepo *TransactionRepo) GetTransactionByDonasiID(donasiID string) (model.Transaction, error) {
	var transaction model.Transaction
	if err := transactionRepo.db.Preload("User").Where("transaction_id = ?", donasiID).First(&transaction).Error; err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}

func (transactionRepo *TransactionRepo) UpdateTransactionStatusByDonasiID(donasiID string, status string) error {
	return transactionRepo.db.Model(&model.Transaction{}).Where("transaction_id = ?", donasiID).Update("status", status).Error
}