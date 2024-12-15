package transaction

import (
	"backend_relawanku/helper"
	"backend_relawanku/model"
	transactionRepo "backend_relawanku/repository/transaction"
	"fmt"
	"time"

	"github.com/midtrans/midtrans-go"
)

func NewTransactionService(tr transactionRepo.TransactionRepository, midtransClient *helper.MidtransClient) *TransactionService {
	return &TransactionService{
		transactionRepoInterface: tr,
		midtransClient: midtransClient,
	}
}

type TransactionService struct {
	transactionRepoInterface transactionRepo.TransactionRepository
	midtransClient *helper.MidtransClient
}

func (s *TransactionService) CreateDonasiTransaction(donasi model.Donasi, user model.User, nominal float64, note string) (string, string, error) {
	// Generate unique order ID
	orderID := fmt.Sprintf("DONASI-%d-%d", donasi.ID, time.Now().UnixNano())

	// Buat transaksi baru
	transaction := &model.Transaction{
		Nominal:   nominal,
		Note:      note,
		DonasiID:  donasi.ID,
		UserID:    user.ID,
	}

	// Siapkan detail transaksi Midtrans
	snapReq := &midtrans.TransactionDetails{
		OrderID:  orderID,
		GrossAmt: int64(nominal),
	}

	// Generate token Midtrans
	snapResp, err := s.midtransClient.GenerateToken(snapReq)
	if err != nil {
		return "", "", err
	}

	// Simpan transaksi ke database
	err = s.transactionRepoInterface.CreateTransaction(transaction)
	if err != nil {
		return "", "", err
	}

	return snapResp.Token, snapResp.RedirectURL, nil
}

func (s *TransactionService) ProcessMidtransCallback(payload map[string]interface{}) error {
	// Verifikasi signature
	if !s.midtransClient.VerifyCallback(payload) {
		return fmt.Errorf("invalid midtrans signature")
	}

	orderID := payload["order_id"].(string)
	status := payload["transaction_status"].(string)

	// Parse order ID untuk mendapatkan transaction ID
	var transactionID uint
	_, err := fmt.Sscanf(orderID, "DONASI-%d-%*d", &transactionID)
	if err != nil {
		return fmt.Errorf("invalid order ID format")
	}

	// Update status transaksi
	return s.transactionRepoInterface.UpdateTransactionStatus(transactionID, status)
}

func (s *TransactionService) GetUserTransactions(userID uint) ([]model.Transaction, error) {
	return s.transactionRepoInterface.GetUserTransactions(userID)
}

func (s *TransactionService) GetDonasiTransactions(donasiID uint) ([]model.Transaction, error) {
	return s.transactionRepoInterface.GetTransactionsByDonasiID(donasiID)
}