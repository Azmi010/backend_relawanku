package transaction

import "backend_relawanku/model"

type TransactionServiceInterface interface {
	CreateDonasiTransaction(userID, donasiID uint, nominal float64) (model.Transaction, error)
}