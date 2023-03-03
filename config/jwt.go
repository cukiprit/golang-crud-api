package config

import (
	"learn/api/models"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"iss": "myapp",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secretkey"))
}
