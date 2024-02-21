package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-rest-template/api/shared/useCases/healthCheck"
	"go-rest-template/internal/app/utils"
	"go-rest-template/internal/app/web"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// TODO: create struct for environment variables and set them elsewhere
	var port = os.Getenv("PORT")
	if port == "" {
		panic("PORT environment variable is required")
	}

	utils.Print("[Server] Running on port", port)

	web.GET("/", healthCheck.UseCase)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), web.LogRequest(http.DefaultServeMux)))
}
