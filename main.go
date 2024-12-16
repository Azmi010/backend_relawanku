package main

import (
	"backend_relawanku/config"

	controllerPro "backend_relawanku/controller/program"
	repoPro "backend_relawanku/repository/program"
	userRepo "backend_relawanku/repository/user"

	articleController "backend_relawanku/controller/article"
	dashboardController "backend_relawanku/controller/dashboard"
	servicePro "backend_relawanku/service/program"
	userService "backend_relawanku/service/user"

	articleRepo "backend_relawanku/repository/article"
	authRepo "backend_relawanku/repository/auth"
	articleService "backend_relawanku/service/article"
	authService "backend_relawanku/service/auth"

	authController "backend_relawanku/controller/auth"
	donasiController "backend_relawanku/controller/donasi"
	registController "backend_relawanku/controller/registration"
	transactionController "backend_relawanku/controller/transaction"
	userController "backend_relawanku/controller/user"
	"backend_relawanku/helper"
	"backend_relawanku/middleware"
	donasiRepo "backend_relawanku/repository/donasi"
	registRepo "backend_relawanku/repository/registration"
	transactionRepo "backend_relawanku/repository/transaction"
	"backend_relawanku/routes"
	donasiService "backend_relawanku/service/donasi"
	registService "backend_relawanku/service/registration"
	transactionService "backend_relawanku/service/transaction"

	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	_ "backend_relawanku/docs"

	cors "github.com/labstack/echo/v4/middleware"
	"github.com/midtrans/midtrans-go"
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

	// Auth setup
	midtransClient := helper.NewMidtransClient(helper.MidtransConfig{
		ServerKey:   os.Getenv("MIDTRANS_SERVER_KEY"),
		ClientKey:   os.Getenv("MIDTRANS_CLIENT_KEY"),
		Environment: midtrans.Sandbox,
	})

	authRepo := authRepo.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	articleRepo := articleRepo.NewArticleRepository(db)
	articleService := articleService.NewArticleService(articleRepo)
	articleController := articleController.NewArticleController(articleService)

	programRepo := repoPro.NewProgramRepository(db)  
	programService := servicePro.NewProgramService(programRepo)  
	programController := controllerPro.NewProgramController(programService)  

	registrationRepo := registRepo.NewUserProgramRepository(db)
	registrationService := registService.NewUserProgramService(registrationRepo)
	registrationController := registController.NewUserProgramController(registrationService)
	userRepo := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	userController := userController.NewUserController(userService)

	donasiRepo := donasiRepo.NewDonasiRepository(db)
	donasiService := donasiService.NewDonasiService(donasiRepo)
	donasiController := donasiController.NewDonasiController(donasiService)

	transactionRepo := transactionRepo.NewTransactionRepository(db)
	transactionService := transactionService.NewTransactionService(transactionRepo, midtransClient)
	transactionController := transactionController.NewTransactionController(transactionService, donasiService, userService)

	dashboardController := dashboardController.NewDashboardController(articleController, programController, donasiController)

	routeController := routes.RouteController{
		AuthController:   authController,
		ProgramController: programController,  
		ArticleController: articleController,
		DashboardController: dashboardController,
		DonasiController:      donasiController,
		TransactionController: transactionController,
		RegisterController:  registrationController,
		UserController: userController,
	}
	routeController.InitRoute(e)

	e.Start(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("failed to load env")
	}
}