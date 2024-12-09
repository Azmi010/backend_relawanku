package routes

import (
	"backend_relawanku/controller/article"
	"backend_relawanku/controller/auth"
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
	
	eJWTAdmin := e.Group("/api/v1/admin", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY_ADMIN")),
	}))
	eJWTAdmin.POST("/article", rc.ArticleController.CreateArticleController)
	eJWTAdmin.GET("/articles", rc.ArticleController.GetAllArticlesController)
	eJWTAdmin.PUT("/article/:id", rc.ArticleController.UpdateArticleController)
	eJWTAdmin.DELETE("/article/:id", rc.ArticleController.DeleteArticleController)

	eJWTUser := e.Group("/api/v1/user", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY_USER")),
	}))
	eJWTUser.GET("/articles", rc.ArticleController.GetAllArticlesController)
	eJWTUser.GET("/articles/category", rc.ArticleController.GetArticlesByCategoryController)
	eJWTUser.GET("/articles/:id", rc.ArticleController.GetArticleByIDController)
}