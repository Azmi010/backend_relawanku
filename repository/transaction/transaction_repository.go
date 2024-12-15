package transaction

import (
	"backend_relawanku/model"

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

func (transactionRepo *TransactionRepo) CreateTransaction(transaction *model.Transaction) error {
	return transactionRepo.db.Create(transaction).Error
}

func (transactionRepo *TransactionRepo) UpdateTransactionStatus(transactionID uint, status string) error {
	return transactionRepo.db.Model(&model.Transaction{}).
		Where("id = ?", transactionID).
		Update("status", status).Error
}

func (transactionRepo *TransactionRepo) GetTransactionsByDonasiID(donasiID uint) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := transactionRepo.db.Where("donasi_id = ?", donasiID).Find(&transactions).Error
	return transactions, err
}

func (transactionRepo *TransactionRepo) GetUserTransactions(userID uint) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := transactionRepo.db.Where("user_id = ?", userID).
		Preload("Donasi").
		Find(&transactions).Error
	return transactions, err
}

func (transactionRepo *TransactionRepo) GetTransactionByID(transactionID uint) (*model.Transaction, error) {
	var transaction model.Transaction
	err := transactionRepo.db.
		Preload("Donasi").
		Preload("User").
		First(&transaction, transactionID).Error
	return &transaction, err
}