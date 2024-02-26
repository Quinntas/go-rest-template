package web

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc[T interface{}] func(response http.ResponseWriter, request *http.Request, data *DecodedRequest[T]) *HttpError

type DecodedRequest[T interface{}] struct {
	Json T
}

func handleRequest[T interface{}](method, path string, handler HandlerFunc[T]) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method != method {
			MethodNotAllowed(response)
			return
		}

		var decodedRequest DecodedRequest[T]

		switch request.Header.Get("Content-Type") {
		case "application/json;charset=UTF-8":
		case "application/json; charset=UTF-8":
		case "application/json":
			err := json.NewDecoder(request.Body).Decode(&decodedRequest.Json)
			if err != nil {
				BadRequest(response)
				return
			}
		}

		err := handler(response, request, &decodedRequest)
		if err != nil {
			_ = HandleHttpError(response, err)
			return
		}
	}
}

func GET[T interface{}](path string, handler HandlerFunc[T]) {
	http.HandleFunc(path, handleRequest[T](http.MethodGet, path, handler))
}

func POST[T interface{}](path string, handler HandlerFunc[T]) {
	http.HandleFunc(path, handleRequest[T](http.MethodPost, path, handler))
}
