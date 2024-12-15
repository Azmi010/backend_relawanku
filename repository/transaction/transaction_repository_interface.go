package transaction

import "backend_relawanku/model"

type TransactionRepository interface {
	CreateTransaction(transaction *model.Transaction) error
	UpdateTransactionStatus(transactionID uint, status string) error
	GetTransactionsByDonasiID(donasiID uint) ([]model.Transaction, error)
	GetUserTransactions(userID uint) ([]model.Transaction, error)
	GetTransactionByID(transactionID uint) (*model.Transaction, error)
}