package middleware

import (
	"backend_relawanku/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAlta struct {
}

type jwtCustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT generates JWT token based on user role
func (jwtAlta JwtAlta) GenerateJWT(username string, role model.UserRole) (string, error) {
	// Define claims
	claims := &jwtCustomClaims{
		Username: username,
		Role:     string(role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // Token expires in 72 hours
		},
	}

	// Select the secret key based on user role
	var secretKey string
	if role == model.RoleAdmin { // Assuming `model.AdminRole` defines the admin role
		secretKey = os.Getenv("JWT_SECRET_KEY_ADMIN")
	} else {
		secretKey = os.Getenv("JWT_SECRET_KEY_USER")
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
