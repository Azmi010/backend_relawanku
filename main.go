package main

import (
	"backend_relawanku/config"
	authController "backend_relawanku/controller/auth"
	donasiController "backend_relawanku/controller/donasi"
	transactionController "backend_relawanku/controller/transaction"
	userController "backend_relawanku/controller/user"
	"backend_relawanku/helper"
	"backend_relawanku/middleware"
	authRepo "backend_relawanku/repository/auth"
	donasiRepo "backend_relawanku/repository/donasi"
	transactionRepo "backend_relawanku/repository/transaction"
	userRepo "backend_relawanku/repository/user"
	"backend_relawanku/routes"
	authService "backend_relawanku/service/auth"
	donasiService "backend_relawanku/service/donasi"
	transactionService "backend_relawanku/service/transaction"
	userService "backend_relawanku/service/user"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	cors "github.com/labstack/echo/v4/middleware"
	_ "backend_relawanku/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           RelawanKu API
// @version         1.0
// @description     API untuk aplikasi RelawanKu
// @host            relawanku.xyz

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	loadEnv()
	db, _ := config.ConnectDatabase()
	config.MigrateDB(db)

	e := echo.New()
	e.Use(cors.CORSWithConfig(cors.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
	}))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	authJwt := middleware.JwtAlta{}

	midtransClient := helper.NewMidtransClient(helper.MidtransConfig{
		ServerKey:   os.Getenv("MIDTRANS_SERVER_KEY"),
		ClientKey:   os.Getenv("MIDTRANS_CLIENT_KEY"),
		Environment: midtrans.Sandbox,
	})

	authRepo := authRepo.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	userRepo := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	userController := userController.NewUserController(userService)

	donasiRepo := donasiRepo.NewDonasiRepository(db)
	donasiService := donasiService.NewDonasiService(donasiRepo)
	donasiController := donasiController.NewDonasiController(donasiService)

	transactionRepo := transactionRepo.NewTransactionRepository(db)
	transactionService := transactionService.NewTransactionService(transactionRepo, midtransClient)
	transactionController := transactionController.NewTransactionController(transactionService, donasiService, userService)

	routeController := routes.RouteController{
		AuthController:        authController,
		DonasiController:      donasiController,
		TransactionController: transactionController,
		UserController: userController,
	}
	routeController.InitRoute(e)

	e.Start(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("failed lod env")
	}
}
