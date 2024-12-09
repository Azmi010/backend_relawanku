package routes

import (
	"backend_relawanku/controller/article"
	"backend_relawanku/controller/auth"
	"backend_relawanku/controller/program"
	"backend_relawanku/controller/dashboard"
	"os"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController *auth.AuthController
	ProgramController *program.ProgramController
	ArticleController *article.ArticleController
	DashboardController *dashboard.DashboardController
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
	eJWTAdmin.POST("/program", rc.ProgramController.CreateProgram)
	eJWTAdmin.GET("/programs", rc.ProgramController.GetAllPrograms)
	eJWTAdmin.GET("/program/category/:category", rc.ProgramController.GetProgramsByCategory) 
	eJWTAdmin.GET("/program/latest", rc.ProgramController.GetLatestProgram) 
	eJWTAdmin.PUT("/program/:id", rc.ProgramController.UpdateProgram)
	eJWTAdmin.DELETE("/program/:id", rc.ProgramController.DeleteProgram)
	
	//dashoard admin
	eJWTAdmin.GET("/dashboard", rc.DashboardController.GetDashboardData)
	
	eJWTUser := e.Group("/api/v1/user", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY_USER")),
	}))
	eJWTUser.GET("/articles", rc.ArticleController.GetAllArticlesController)
	eJWTUser.GET("/articles/category", rc.ArticleController.GetArticlesByCategoryController)
	eJWTUser.GET("/articles/:id", rc.ArticleController.GetArticleByIDController)
	eJWTUser.GET("/programs", rc.ProgramController.GetAllPrograms)
	eJWTUser.GET("/program/:id", rc.ProgramController.GetProgramByID)
	eJWTUser.GET("/program/category/:category", rc.ProgramController.GetProgramsByCategory) 
	eJWTUser.GET("/program/latest", rc.ProgramController.GetLatestProgram)

	//beranda user
	eJWTUser.GET("/homePage", rc.ArticleController.GetAllArticlesController)
	eJWTUser.GET("/homePage/:id", rc.ArticleController.GetArticleByIDController)

}