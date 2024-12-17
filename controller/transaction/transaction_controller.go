package transaction

import (
	"backend_relawanku/controller/base"
	"backend_relawanku/controller/transaction/request"
	"backend_relawanku/service/transaction"

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
		// return base.ErrorResponse(c, err, map[string]string{
		// 	"error": "Failed to bind request",
		// })
	}

	transaction, _ := req.TransactionToModel()

	created, err := transactionController.transactionServiceInterfae.CreateDonasiTransaction(transaction.UserID, transaction.DonasiID, transaction.Nominal)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, created)
}
