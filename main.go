package main

import (
	"backend_relawanku/config"
	authController "backend_relawanku/controller/auth"
	donasiController "backend_relawanku/controller/donasi"
	"backend_relawanku/middleware"
	authRepo "backend_relawanku/repository/auth"
	donasiRepo "backend_relawanku/repository/donasi"
	"backend_relawanku/routes"
	authService "backend_relawanku/service/auth"
	donasiService "backend_relawanku/service/donasi"
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

	authRepo := authRepo.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	donasiRepo := donasiRepo.NewDonasiRepository(db)
	donasiService := donasiService.NewDonasiService(donasiRepo)
	donasiController := donasiController.NewDonasiController(donasiService)

	routeController := routes.RouteController{
		AuthController:   authController,
		DonasiController: donasiController,
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