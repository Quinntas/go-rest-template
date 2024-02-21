package utils

import (
	"github.com/joho/godotenv"
	"os"
)

type EnvVariables struct {
	Port        string
	DatabaseURL string
}

var Env *EnvVariables

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	Env = &EnvVariables{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
