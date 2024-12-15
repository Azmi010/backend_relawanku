package transaction

import (
	"backend_relawanku/service/donasi"
	"backend_relawanku/service/transaction"
	"backend_relawanku/service/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewTransactionController(tc transaction.TransactionServiceInterface, dc donasi.DonasiServiceInterface,
	uc user.UserServiceInterface) *TransactionController {
	return &TransactionController{
		transactionServiceInterfae: tc,
		donasiServiceInterface: dc,
		userServiceInterface: uc,
	}
}

type TransactionController struct {
	transactionServiceInterfae transaction.TransactionServiceInterface
	donasiServiceInterface donasi.DonasiServiceInterface
	userServiceInterface user.UserServiceInterface
}

func (transactionController *TransactionController) CreateTransactionController(c echo.Context) error {
	// Dapatkan data dari request
	donasiID, _ := strconv.ParseUint(c.FormValue("donasi_id"), 10, 64)
	nominal, _ := strconv.ParseFloat(c.FormValue("nominal"), 64)
	note := c.FormValue("note")

	// Dapatkan user dari session/token
	userID := c.Get("user_id").(uint)
	user, err := transactionController.userServiceInterface.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	// Dapatkan donasi
	donasi, err := transactionController.donasiServiceInterface.GetDonasiById(uint(donasiID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Donasi not found"})
	}

	// Buat transaksi
	token, paymentURL, err := transactionController.transactionServiceInterfae.CreateDonasiTransaction(
		donasi, 
		user, 
		nominal, 
		note,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token":       token,
		"payment_url": paymentURL,
	})
}

func (transactionController *TransactionController) MidtransCallback(c echo.Context) error {
	var payload map[string]interface{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := transactionController.transactionServiceInterfae.ProcessMidtransCallback(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

func (transactionController *TransactionController) GetUserTransactions(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	transactions, err := transactionController.transactionServiceInterfae.GetUserTransactions(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, transactions)
}

func (transactionController *TransactionController) GetDonasiTransactions(c echo.Context) error {
	donasiID, _ := strconv.ParseUint(c.Param("donasi_id"), 10, 64)

	transactions, err := transactionController.transactionServiceInterfae.GetDonasiTransactions(uint(donasiID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, transactions)
}