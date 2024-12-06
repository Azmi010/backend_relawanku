package routes

import (
	"backend_relawanku/controller/auth"
	"backend_relawanku/controller/program"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController *auth.AuthController
	ProgramController *program.ProgramController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	e.POST("/api/v1/register", rc.AuthController.RegisterController)
	e.POST("/api/v1/login", rc.AuthController.LoginController)

	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	eAdmin := eJWT.Group("/admin")
	eAdmin.POST("/api/v1/program", rc.ProgramController.CreateProgram)
	eAdmin.GET("/api/v1/programs", rc.ProgramController.GetAllPrograms)
	eAdmin.GET("/api/v1/program/category/:category", rc.ProgramController.GetProgramsByCategory) 
	eAdmin.GET("/api/v1/program/latest", rc.ProgramController.GetLatestProgram) 
	eAdmin.PUT("/api/v1/program/:id", rc.ProgramController.UpdateProgram)
	eAdmin.DELETE("/api/v1/program/:id", rc.ProgramController.DeleteProgram)

	eUser := eJWT.Group("/user")
	eUser.GET("/api/v1/programs", rc.ProgramController.GetAllPrograms)
	eUser.GET("/api/v1/program/:id", rc.ProgramController.GetProgramByID)
	eUser.GET("/api/v1/program/category/:category", rc.ProgramController.GetProgramsByCategory) 
	eUser.GET("/api/v1/program/latest", rc.ProgramController.GetLatestProgram) 

}