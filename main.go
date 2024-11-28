package main

import (
	"backend_relawanku/config"
	controller "backend_relawanku/controller/auth"
	repo "backend_relawanku/repository/auth"
	"backend_relawanku/routes"
	service "backend_relawanku/service/auth"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	db, _ := config.ConnectDatabase()
	config.MigrateDB(db)

	e := echo.New()

	authRepo := repo.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	routeController := routes.RouteController{
		AuthController:   authController,
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