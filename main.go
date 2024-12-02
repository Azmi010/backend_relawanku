package main

import (
	"backend_relawanku/config"
	controller "backend_relawanku/controller/auth"
	controllerPro "backend_relawanku/controller/program" 
	"backend_relawanku/middleware"
	repo "backend_relawanku/repository/auth"
	repoPro "backend_relawanku/repository/program" 
	"backend_relawanku/routes"
	service "backend_relawanku/service/auth"
	servicePro "backend_relawanku/service/program" 
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
	authRepo := repo.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo, authJwt)
	authController := controller.NewAuthController(authService)

	programRepo := repoPro.NewProgramRepository(db)  
	programService := servicePro.NewProgramService(programRepo)  
	programController := controllerPro.NewProgramController(programService)  

	routeController := routes.RouteController{
		AuthController:   authController,
		ProgramController: programController,  
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
