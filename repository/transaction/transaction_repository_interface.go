package transaction

import "backend_relawanku/model"

type TransactionRepository interface {
	CreateTransaction(transaction model.Transaction) (model.Transaction, error)
	GetTransactionByID(transactionID int) (model.Transaction, error)
	GetAllTransactions() ([]model.Transaction, error)
	UpdateTransaction(transactionID int, updates map[string]interface{}) (model.Transaction, error)
	UpdateTransactionStatus(transactionID int, status string) error
	DeleteTransaction(transactionID int) error
	GetTransactionByDonasiID(donasiID string) (model.Transaction, error)
	UpdateTransactionStatusByDonasiID(donasiID string, status string) error
}