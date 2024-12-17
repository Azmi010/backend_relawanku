package transaction

import (
	"backend_relawanku/helper"
	"backend_relawanku/model"
	donasiRepo "backend_relawanku/repository/donasi"
	transactionRepo "backend_relawanku/repository/transaction"
	"errors"
	"fmt"
)

func NewTransactionService(tr transactionRepo.TransactionRepository, dr donasiRepo.DonasiRepository) *TransactionService {
	return &TransactionService{
		transactionRepoInterface: tr,
		donasiRepoInterface:      dr,
	}
}

type TransactionService struct {
	transactionRepoInterface transactionRepo.TransactionRepository
	donasiRepoInterface      donasiRepo.DonasiRepository
}

func (s *TransactionService) CreateDonasiTransaction(userID, donasiID uint, nominal float64) (model.Transaction, error) {
	if userID == 0 || donasiID == 0 {
		return model.Transaction{}, errors.New("user_id or donasi_id is missing")
	}

	donasi, err := s.donasiRepoInterface.GetDonasiById(donasiID)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("failed to fetch donasi: %w", err)
	}
	
	transaction := model.Transaction{
		UserID:        userID,
		DonasiID:      donasiID,
		Nominal:       nominal,
		Status:        "Pending",
		TransactionID: "ORDER-" + helper.GenerateUniqueID(),
	}

	paymentUrl, err := helper.CreateTransaction(
		transaction.TransactionID,
		int64(nominal),
		donasi.Title,
		donasi.Category,
		donasi.Category,
		donasi.Category,
	)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("failed to create payment URL: %w", err)
	}
	transaction.PaymentUrl = paymentUrl

	createdTransaction, err := s.transactionRepoInterface.CreateTransaction(transaction)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("failed to save transaction: %w", err)
	}

	return createdTransaction, nil
}