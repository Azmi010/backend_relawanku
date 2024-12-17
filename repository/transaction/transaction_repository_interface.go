package transaction

import "backend_relawanku/model"

type TransactionRepository interface {
	CreateTransaction(transaction model.Transaction) (model.Transaction, error)
}