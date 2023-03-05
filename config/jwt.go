package config

import (
	"learn/api/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user models.User) (string, error) {
	// Create a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte("secretkey"))

	return tokenString, err
}
