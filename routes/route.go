package routes

import (
	"backend_relawanku/controller/article"
	"backend_relawanku/controller/auth"
	"os"
	"backend_relawanku/controller/program"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController *auth.AuthController
	ProgramController *program.ProgramController
	ArticleController *article.ArticleController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	e.POST("/api/v1/register", rc.AuthController.RegisterController)
	e.POST("/api/v1/login", rc.AuthController.LoginController)
	
	
	eJWTAdmin := e.Group("/admin", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY_ADMIN")),
	}))
	eJWTAdmin.POST("/api/v1/article", rc.ArticleController.CreateArticleController)
	eJWTAdmin.GET("/api/v1/articles", rc.ArticleController.GetAllArticlesController)
	eJWTAdmin.PUT("/api/v1/article/:id", rc.ArticleController.UpdateArticleController)
	eJWTAdmin.DELETE("/api/v1/article/:id", rc.ArticleController.DeleteArticleController)
	eJWTAdmin.POST("/api/v1/program", rc.ProgramController.CreateProgram)
	eJWTAdmin.GET("/api/v1/programs", rc.ProgramController.GetAllPrograms)
	eJWTAdmin.GET("/api/v1/program/category/:category", rc.ProgramController.GetProgramsByCategory) 
	eJWTAdmin.GET("/api/v1/program/latest", rc.ProgramController.GetLatestProgram) 
	eJWTAdmin.PUT("/api/v1/program/:id", rc.ProgramController.UpdateProgram)
	eJWTAdmin.DELETE("/api/v1/program/:id", rc.ProgramController.DeleteProgram)

	eJWTUser := e.Group("/user", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY_USER")),
	}))
	eJWTUser.GET("/api/v1/articles", rc.ArticleController.GetAllArticlesController)
	eJWTUser.GET("/api/v1/articles/category", rc.ArticleController.GetArticlesByCategoryController)
	eJWTUser.GET("/api/v1/articles/:id", rc.ArticleController.GetArticleByIDController)
	eJWTUser.GET("/api/v1/programs", rc.ProgramController.GetAllPrograms)
	eJWTUser.GET("/api/v1/program/:id", rc.ProgramController.GetProgramByID)
	eJWTUser.GET("/api/v1/program/category/:category", rc.ProgramController.GetProgramsByCategory) 
	eJWTUser.GET("/api/v1/program/latest", rc.ProgramController.GetLatestProgram)
	
	//ini untuk beranda (nampilin artikel)
	eJWTUser.GET("/api/v1/homePage", rc.ArticleController.GetAllArticlesController)
	eJWTUser.GET("/api/v1/homePage/:id", rc.ArticleController.GetArticleByIDController)

}