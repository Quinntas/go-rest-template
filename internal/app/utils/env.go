package utils

import (
	"github.com/joho/godotenv"
	"os"
)

type EnvVariables struct {
	Port        string
	DatabaseURL string
	PEPPER      string
	JwtSecret   string
	RedisURL    string
}

var Env *EnvVariables

func getEnv(key string, required bool) string {
	value := os.Getenv(key)
	if required && value == "" {
		panic("Missing required environment variable: " + key)
	}
	return value
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	Env = &EnvVariables{
		Port:        getEnv("PORT", true),
		DatabaseURL: getEnv("DATABASE_URL", true),
		PEPPER:      getEnv("PEPPER", true),
		JwtSecret:   getEnv("JWT_SECRET", true),
		RedisURL:    getEnv("REDIS_URL", true),
	}
}
