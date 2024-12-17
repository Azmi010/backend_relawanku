package request

import "backend_relawanku/model"

type TransactionRequest struct {
	UserID   uint `json:"user_id"`
	DonasiID uint `json:"donasi_id"`
	Nominal uint `json:"nominal"`
}

func (transactionRequest TransactionRequest) TransactionToModel() (model.Transaction, error) {
	transaction := model.Transaction{
		UserID: transactionRequest.UserID,
		DonasiID: transactionRequest.DonasiID,
		Nominal: float64(transactionRequest.Nominal),
	}

	return transaction, nil
}