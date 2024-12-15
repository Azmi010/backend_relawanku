package routes

import (
	"backend_relawanku/controller/article"
	"backend_relawanku/controller/auth"
	"backend_relawanku/controller/dashboard"
	"backend_relawanku/controller/program"
	"backend_relawanku/controller/registration"
	"backend_relawanku/controller/user"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"backend_relawanku/controller/donasi"
	"backend_relawanku/controller/transaction"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController *auth.AuthController
	ProgramController *program.ProgramController
	ArticleController *article.ArticleController
	DashboardController *dashboard.DashboardController
	DonasiController *donasi.DonasiController
	TransactionController *transaction.TransactionController
	RegisterController *registration.UserProgramController
	UserController *user.UserController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	e.POST("/api/v1/register", rc.AuthController.RegisterController)
	e.POST("/api/v1/login", rc.AuthController.LoginController)
	
	e.POST("/midtrans-callback", rc.TransactionController.MidtransCallback)
	
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
	eJWTAdmin.POST("/donasi", rc.DonasiController.CreateDonasiController)
	eJWTAdmin.PUT("/donasi/:id", rc.DonasiController.UpdateDonasiController)
	eJWTAdmin.DELETE("/donasi/:id", rc.DonasiController.DeleteDonasiController)
	eJWTAdmin.GET("/donasi", rc.DonasiController.GetAllDonasiController)
	
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
	eJWTUser.POST("/register-program", rc.RegisterController.RegisterProgram)
	eJWTUser.GET("/my-program/:id", rc.RegisterController.GetUserPrograms)
	eJWTUser.GET("/donasi", rc.DonasiController.GetAllDonasiController)
	eJWTUser.GET("/donasi/:category", rc.DonasiController.GetDonasiByCategoryController)
	eJWTUser.GET("/donasi/:id", rc.DonasiController.GetDonasiByIdController)
	eJWTUser.POST("/transaction", rc.TransactionController.CreateTransactionController)
	eJWTUser.GET("/transactions", rc.TransactionController.GetUserTransactions)
	eJWTUser.GET("/donasi/:id/transactions", rc.TransactionController.GetDonasiTransactions)

	//beranda user
	eJWTUser.GET("/homePage", rc.ArticleController.GetAllArticlesController)
	eJWTUser.GET("/homePage/:id", rc.ArticleController.GetArticleByIDController)

	// profile user
	eJWTUser.GET("/profile/:id", rc.UserController.GetUserByIDController)
	eJWTUser.PUT("/profile/:id", rc.UserController.UpdateUserController)
	eJWTUser.PUT("/profile/:id", rc.UserController.UpdatePasswordController)
}