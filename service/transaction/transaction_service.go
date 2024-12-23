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
		DonasiID:      donasi.ID,
		Nominal:       nominal,
		Status:        "Pending",
		TransactionID: "ORDER-" + helper.GenerateUniqueID(),
	}

	paymentUrl, err := helper.CreateTransaction(
		transaction.TransactionID,
		int64(nominal),
		transaction.User.Username,
		transaction.User.Email,
		transaction.User.Gender,
		transaction.User.Address,
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

func (s *TransactionService) HandleMidtransNotification(donasiID, status string) error {
	_, err := s.transactionRepoInterface.GetTransactionByDonasiID(donasiID)
	if err != nil {
		return fmt.Errorf("transaction not found for order ID %s: %w", donasiID, err)
	}

	var updatedStatus string
	switch status {
	case "capture", "settlement":
		updatedStatus = "complete"
	case "pending":
		updatedStatus = "pending"
	case "deny", "expire", "cancel":
		updatedStatus = "failed"
	default:
		updatedStatus = "unknown"
	}

	if err := s.transactionRepoInterface.UpdateTransactionStatusByDonasiID(donasiID, updatedStatus); err != nil {
		return fmt.Errorf("failed to update transaction status: %w", err)
	}

	return nil
}

func (s *TransactionService) GetTransactionByID(transactionID int) (model.Transaction, error) {
	return s.transactionRepoInterface.GetTransactionByID(transactionID)
}

func (s *TransactionService) GetAllTransactions() ([]model.Transaction, error) {
	return s.transactionRepoInterface.GetAllTransactions()
}

func (s *TransactionService) UpdateTransaction(id int, updatedTransaction model.Transaction) (model.Transaction, error) {
	if id <= 0 {
		return model.Transaction{}, errors.New("invalid transaction ID")
	}

	updates := map[string]interface{}{
		"user_id":     updatedTransaction.UserID,
		"total_price": updatedTransaction.Nominal + 2000,
		"status":      updatedTransaction.Status,
		"payment_url": updatedTransaction.PaymentUrl,
		"updated_at":  helper.GetCurrentTime(),
	}

	return s.transactionRepoInterface.UpdateTransaction(id, updates)
}

func (s *TransactionService) UpdateTransactionStatus(transactionID int, status string) error {
	if status == "" {
		return errors.New("status cannot be empty")
	}
	return s.transactionRepoInterface.UpdateTransactionStatus(transactionID, status)
}

func (s *TransactionService) DeleteTransaction(transactionID int) error {
	return s.transactionRepoInterface.DeleteTransaction(transactionID)
}