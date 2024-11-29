package middleware

import (
	"backend_relawanku/controller/base"
	"backend_relawanku/model"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func RoleMiddleware(allowedRoles ...model.UserRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := c.Get("user").(jwt.MapClaims)
			role := model.UserRole(claims["role"].(string))

			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, base.ErrorResponse(c, errors.New("access denied")))
		}
	}
}