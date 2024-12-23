package transaction

import "backend_relawanku/model"

type TransactionServiceInterface interface {
	CreateDonasiTransaction(userID, donasiID uint, nominal float64) (model.Transaction, error)
	HandleMidtransNotification(donasiID, status string) error
	GetTransactionByID(transactionID int) (model.Transaction, error)
	GetAllTransactions() ([]model.Transaction, error)
	UpdateTransaction(id int, updatedTransaction model.Transaction) (model.Transaction, error)
	UpdateTransactionStatus(transactionID int, status string) error
	DeleteTransaction(transactionID int) error
}