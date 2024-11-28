package routes

import (
	"backend_relawanku/controller/auth"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController *auth.AuthController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	e.POST("/api/v1/register", rc.AuthController.RegisterController)
}