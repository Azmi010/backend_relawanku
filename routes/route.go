package routes

import (
	"backend_relawanku/controller/article"
	"backend_relawanku/controller/auth"
	// "backend_relawanku/middleware"
	// "backend_relawanku/model"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController    *auth.AuthController
	ArticleController *article.ArticleController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	e.POST("/api/v1/register", rc.AuthController.RegisterController)
	e.POST("/api/v1/login", rc.AuthController.LoginController)
	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	eAdmin := eJWT.Group("/admin")
	eAdmin.POST("/api/v1/article", rc.ArticleController.CreateArticleController)

	eUser := eJWT.Group("/user")
	eUser.GET("/api/v1/articles", rc.ArticleController.GetAllArticlesController)
	eUser.GET("/api/v1/articles/category", rc.ArticleController.GetArticlesByCategoryController)
	eUser.GET("/api/v1/articles/:id", rc.ArticleController.GetArticleByIDController)
}