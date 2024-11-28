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

func (jwtAlta JwtAlta) GenerateJWT(username string, role model.UserRole) (string, error) {
	claims := &jwtCustomClaims{
		username,
		string(role),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}
