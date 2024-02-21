package main

import (
	"fmt"
	"github.com/joho/godotenv"
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

	var port = os.Getenv("PORT")

	if port == "" {
		panic("PORT environment variable is required")
	}

	utils.Print("[Server] Running on port", port)

	web.GET("/", func(response http.ResponseWriter, request *http.Request, data web.DecodedRequest[struct{}]) {
		web.JsonResponse(response, http.StatusOK, map[string]interface{}{
			"message": "Hello, world!",
		})
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), web.LogRequest(http.DefaultServeMux)))
}
