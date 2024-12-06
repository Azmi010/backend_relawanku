package routes

import (
	"backend_relawanku/controller/auth"
	"backend_relawanku/controller/donasi"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController *auth.AuthController
	DonasiController *donasi.DonasiController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	e.POST("/api/v1/register", rc.AuthController.RegisterController)
	e.POST("/api/v1/login", rc.AuthController.LoginController)
	
	eJWTAdmin := e.Group("/admin", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY_ADMIN")),
	}))
	eJWTAdmin.POST("/api/v1/donasi", rc.DonasiController.CreateDonasiController)
	eJWTAdmin.PUT("/api/v1/donasi/:id", rc.DonasiController.UpdateDonasiController)
	eJWTAdmin.DELETE("/api/v1/donasi/:id", rc.DonasiController.DeleteDonasiController)
	eJWTAdmin.GET("/api/v1/donasi", rc.DonasiController.GetAllDonasiController)

	eJWTUser := e.Group("/user", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY_USER")),
	}))
	eJWTUser.GET("/api/v1/donasi", rc.DonasiController.GetAllDonasiController)
	eJWTUser.GET("/api/v1/donasi/:category", rc.DonasiController.GetDonasiByCategoryController)
	eJWTUser.GET("/api/v1/donasi/:id", rc.DonasiController.GetDonasiByIdController)
}