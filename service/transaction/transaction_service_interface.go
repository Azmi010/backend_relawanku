package transaction

import "backend_relawanku/model"

type TransactionServiceInterface interface {
	CreateDonasiTransaction(donasi model.Donasi, user model.User, nominal float64, note string,) (string, string, error)
	ProcessMidtransCallback(payload map[string]interface{}) error
	GetUserTransactions(userID uint) ([]model.Transaction, error)
	GetDonasiTransactions(donasiID uint) ([]model.Transaction, error)
}