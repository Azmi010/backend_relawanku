package main

import (
	"backend_relawanku/config"

	controllerPro "backend_relawanku/controller/program"
	repoPro "backend_relawanku/repository/program"
	userRepo "backend_relawanku/repository/user"

	"backend_relawanku/routes"

	articleController "backend_relawanku/controller/article"
	authController "backend_relawanku/controller/auth"
	dashboardController "backend_relawanku/controller/dashboard"
	userController "backend_relawanku/controller/user"
	servicePro "backend_relawanku/service/program"
	userService "backend_relawanku/service/user"

	"backend_relawanku/middleware"

	articleRepo "backend_relawanku/repository/article"
	authRepo "backend_relawanku/repository/auth"
	articleService "backend_relawanku/service/article"
	authService "backend_relawanku/service/auth"

	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	// Auth setup
	authRepo := authRepo.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	articleRepo := articleRepo.NewArticleRepository(db)
	articleService := articleService.NewArticleService(articleRepo)
	articleController := articleController.NewArticleController(articleService)

	programRepo := repoPro.NewProgramRepository(db)  
	programService := servicePro.NewProgramService(programRepo)  
	programController := controllerPro.NewProgramController(programService)  

	dashboardController := dashboardController.NewDashboardController(articleController, programController)

	userRepo := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	userController := userController.NewUserController(userService)

	routeController := routes.RouteController{
		AuthController:   authController,
		ProgramController: programController,  
		ArticleController: articleController,
		DashboardController: dashboardController,
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
