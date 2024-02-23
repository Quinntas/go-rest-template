package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJsonWebToken[T interface{} | string](data *T, expirationTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"data": *data,
			"exp":  expirationTime,
		})
	tokenString, err := token.SignedString([]byte(Env.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
