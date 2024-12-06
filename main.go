package main

import (
	"backend_relawanku/config"

	controllerPro "backend_relawanku/controller/program" 
	repoPro "backend_relawanku/repository/program" 

	"backend_relawanku/routes"

	servicePro "backend_relawanku/service/program" 
	authController "backend_relawanku/controller/auth"
	articleController "backend_relawanku/controller/article"

	"backend_relawanku/middleware"

	authRepo "backend_relawanku/repository/auth"
	articleRepo "backend_relawanku/repository/article"
	authService "backend_relawanku/service/auth"
	articleService "backend_relawanku/service/article"
	
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	db, _ := config.ConnectDatabase()
	config.MigrateDB(db)

	e := echo.New()
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

	routeController := routes.RouteController{
		AuthController:   authController,
		ProgramController: programController,  
		ArticleController: articleController,
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
