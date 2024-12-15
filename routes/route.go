package routes

import (
	"backend_relawanku/controller/auth"
	"backend_relawanku/controller/donasi"
	"backend_relawanku/controller/transaction"
	"backend_relawanku/controller/user"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController *auth.AuthController
	DonasiController *donasi.DonasiController
	TransactionController *transaction.TransactionController
	UserController *user.UserController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	e.POST("/api/v1/register", rc.AuthController.RegisterController)
	e.POST("/api/v1/login", rc.AuthController.LoginController)
	e.POST("/midtrans-callback", rc.TransactionController.MidtransCallback)
	
	eJWTAdmin := e.Group("/api/v1/admin", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY_ADMIN")),
	}))
	eJWTAdmin.POST("/donasi", rc.DonasiController.CreateDonasiController)
	eJWTAdmin.PUT("/donasi/:id", rc.DonasiController.UpdateDonasiController)
	eJWTAdmin.DELETE("/donasi/:id", rc.DonasiController.DeleteDonasiController)
	eJWTAdmin.GET("/donasi", rc.DonasiController.GetAllDonasiController)

	eJWTUser := e.Group("/api/v1/user", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY_USER")),
	}))
	eJWTUser.GET("/donasi", rc.DonasiController.GetAllDonasiController)
	eJWTUser.GET("/donasi/:category", rc.DonasiController.GetDonasiByCategoryController)
	eJWTUser.GET("/donasi/:id", rc.DonasiController.GetDonasiByIdController)
	eJWTUser.POST("/transaction", rc.TransactionController.CreateTransactionController)
	eJWTUser.GET("/transactions", rc.TransactionController.GetUserTransactions)
	eJWTUser.GET("/donasi/:id/transactions", rc.TransactionController.GetDonasiTransactions)
}