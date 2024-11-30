package middleware

import "backend_relawanku/model"

type JwtInterface interface {
	GenerateJWT(username string, role model.UserRole) (string, error)
}