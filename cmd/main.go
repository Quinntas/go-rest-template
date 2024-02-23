package main

import (
	"fmt"
	"go-rest-template/api/shared/useCases/healthCheck"
	"go-rest-template/api/user/useCases/createUser"
	"go-rest-template/api/user/useCases/login"
	"go-rest-template/internal/app/database"
	"go-rest-template/internal/app/redis"
	"go-rest-template/internal/app/utils"
	"go-rest-template/internal/app/web"
	"log"
	"net/http"
)

func main() {
	utils.LoadEnv()

	database.ConnectDB(utils.Env.DatabaseURL)
	redis.ConnectRedis(utils.Env.RedisURL)

	utils.Print("[Server] Running on port", utils.Env.Port)

	web.GET("/", healthCheck.UseCase)
	web.POST("/api/v1/users", createUser.UseCase)
	web.POST("/api/v1/users/login", login.UseCase)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", utils.Env.Port), web.LogRequest(http.DefaultServeMux)))
}
