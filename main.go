package main

import (
	"backend_relawanku/config"
	authController "backend_relawanku/controller/auth"
	articleController "backend_relawanku/controller/article"
	"backend_relawanku/middleware"
	authRepo "backend_relawanku/repository/auth"
	articleRepo "backend_relawanku/repository/article"
	"backend_relawanku/routes"
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

	authRepo := authRepo.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	articleRepo := articleRepo.NewArticleRepository(db)
	articleService := articleService.NewArticleService(articleRepo)
	articleController := articleController.NewArticleController(articleService)

	routeController := routes.RouteController{
		AuthController:   authController,
		ArticleController: articleController,
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