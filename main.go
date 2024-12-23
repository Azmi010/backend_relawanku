package main

import (
	// controllers
	articleController "backend_relawanku/controller/article"
	authController "backend_relawanku/controller/auth"
	dashboardController "backend_relawanku/controller/dashboard"
	donasiController "backend_relawanku/controller/donasi"
	programController "backend_relawanku/controller/program"
	registController "backend_relawanku/controller/registration"
	transactionController "backend_relawanku/controller/transaction"
	userController "backend_relawanku/controller/user"
	
	// services
	articleService "backend_relawanku/service/article"
	authService "backend_relawanku/service/auth"
	donasiService "backend_relawanku/service/donasi"
	programService "backend_relawanku/service/program"
	registService "backend_relawanku/service/registration"
	transactionService "backend_relawanku/service/transaction"
	userService "backend_relawanku/service/user"
	
	// repositories
	articleRepo "backend_relawanku/repository/article"
	authRepo "backend_relawanku/repository/auth"
	donasiRepo "backend_relawanku/repository/donasi"
	programRepo "backend_relawanku/repository/program"
	registRepo "backend_relawanku/repository/registration"
	transactionRepo "backend_relawanku/repository/transaction"
	userRepo "backend_relawanku/repository/user"
	
	"backend_relawanku/config"
	"backend_relawanku/middleware"
	"backend_relawanku/routes"
	"backend_relawanku/helper"

	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	cors "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "backend_relawanku/docs"
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
	helper.InitMidtrans()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	authJwt := middleware.JwtAlta{}

	authRepo := authRepo.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	articleRepo := articleRepo.NewArticleRepository(db)
	articleService := articleService.NewArticleService(articleRepo)
	articleController := articleController.NewArticleController(articleService)

	programRepo := programRepo.NewProgramRepository(db)  
	programService := programService.NewProgramService(programRepo)  
	programController := programController.NewProgramController(programService)  

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
	transactionService := transactionService.NewTransactionService(transactionRepo, donasiRepo)
	transactionController := transactionController.NewTransactionController(transactionService)

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