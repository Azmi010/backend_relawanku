package response

import (
	"backend_relawanku/model"
	"time"
)

type TransactionResponse struct {
	ID            int          `json:"id"`
	TransactionID string       `json:"transaction_id"`
	User          UserResponse `json:"user"`
	DonasiID      int          `json:"donasi_id"`
	TotalPrice    float64      `json:"total_price"`
	Status        string       `json:"status"`
	PaymentURL    string       `json:"payment_url"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type UserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

func TransactionFromModel(transaction model.Transaction) TransactionResponse {

	userResponse := UserResponse{
		ID:          int(transaction.User.ID),
		Name:        transaction.User.Username,
		Email:       transaction.User.Email,
		Address:     transaction.User.Address,
	}

	return TransactionResponse{
		ID:            int(transaction.ID),
		TransactionID: transaction.TransactionID,
		User:          userResponse,
		TotalPrice:    transaction.Nominal,
		Status:        transaction.Status,
		PaymentURL:    transaction.PaymentUrl,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}