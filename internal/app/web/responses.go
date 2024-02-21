package web

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(response http.ResponseWriter, status int, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	encoder := json.NewEncoder(response)
	if err := encoder.Encode(data); err != nil {
		// TODO
		return
	}
}

func TextResponse(response http.ResponseWriter, status int, data string) {
	response.Header().Set("Content-Type", "text/plain")
	response.WriteHeader(status)
	_, err := response.Write([]byte(data))
	if err != nil {
		// TODO
		return
	}
}
