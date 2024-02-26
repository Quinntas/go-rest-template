package utils

import (
	"encoding/json"
	"fmt"
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

// TODO: FIX THIS
func DecodedJsonWebToken[T interface{}](tokenString string) (*T, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Env.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	data, ok := claims["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("data is not a map[string]interface{}")
	}
	var result T
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
