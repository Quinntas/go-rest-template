package web

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc[T interface{}] func(response http.ResponseWriter, request *http.Request, data *DecodedRequest[T])

type MiddlewareFunc[T interface{}] func(HandlerFunc[T]) HandlerFunc[T]

type DecodedRequest[T interface{}] struct {
	Json T
}

func handleRequest[T interface{}](method, path string, handler HandlerFunc[T], middlewares ...MiddlewareFunc[T]) http.HandlerFunc {
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

		finalHandler := handler
		for _, middleware := range middlewares {
			finalHandler = middleware(finalHandler)
		}

		finalHandler(response, request, &decodedRequest)
	}
}

func GET[T interface{}](path string, handler HandlerFunc[T], middlewares ...MiddlewareFunc[T]) {
	http.HandleFunc(path, handleRequest[T](http.MethodGet, path, handler, middlewares...))
}

func POST[T interface{}](path string, handler HandlerFunc[T], middlewares ...MiddlewareFunc[T]) {
	http.HandleFunc(path, handleRequest[T](http.MethodPost, path, handler, middlewares...))
}
