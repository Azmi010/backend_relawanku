package transaction

import (
	"backend_relawanku/controller/base"
	"backend_relawanku/controller/transaction/request"
	"backend_relawanku/controller/transaction/response"
	"backend_relawanku/service/transaction"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewTransactionController(tc transaction.TransactionServiceInterface) *TransactionController {
	return &TransactionController{
		transactionServiceInterfae: tc,
	}
}

type TransactionController struct {
	transactionServiceInterfae transaction.TransactionServiceInterface
}

func (transactionController *TransactionController) CreateTransactionController(c echo.Context) error {
	req := new(request.TransactionRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err)
	}

	transaction, _ := req.TransactionToModel()

	created, err := transactionController.transactionServiceInterfae.CreateDonasiTransaction(transaction.UserID, transaction.DonasiID, transaction.Nominal)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.TransactionFromModel(created))
}

func (transactionController *TransactionController) GetTransactionByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := transactionController.transactionServiceInterfae.GetTransactionByID(id)
	if err != nil {
		return base.ErrorResponse(c, errors.New("failed to get transaction"))
		}
	return base.SuccessResponse(c, transaction)
}

func (transactionController *TransactionController) GetAllTransactions(c echo.Context) error {
	transactions, err := transactionController.transactionServiceInterfae.GetAllTransactions()
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	transactionResponses := make([]response.TransactionResponse, 0)
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, response.TransactionFromModel(transaction))
	}

	return base.SuccessResponse(c, transactionResponses)
}

func (transactionController *TransactionController) UpdateTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request.TransactionRequest)

	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err)
	}

	updatedTransaction, err := req.TransactionToModel()
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	transaction, err := transactionController.transactionServiceInterfae.UpdateTransaction(id, updatedTransaction)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.TransactionFromModel(transaction))
}

func (transactionController *TransactionController) UpdateTransactionStatus(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	status := c.Param("status")
	err := transactionController.transactionServiceInterfae.UpdateTransactionStatus(id, status)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, map[string]string{
		"transaction_id": strconv.Itoa(id),
		"status":         status,
	})
}

func (transactionController *TransactionController) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := transactionController.transactionServiceInterfae.DeleteTransaction(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, "Transaction deleted successfully")
}

func (transactionController *TransactionController) HandleMidtransNotification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid notification payload",
		})
	}

	orderID, ok := notificationPayload["donasi_id"].(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Donasi ID is missing",
		})
	}

	transactionStatus, ok := notificationPayload["transaction_status"].(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Transaction status is missing",
		})
	}

	// Update transaction status
	if err := transactionController.transactionServiceInterfae.HandleMidtransNotification(orderID, transactionStatus); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update transaction status",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Notification received successfully",
	})
}