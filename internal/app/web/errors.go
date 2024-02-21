package web

import (
	"net/http"
)

type HttpError struct {
	Code int
	Body interface{}
}

func HandleHttpError(response http.ResponseWriter, err HttpError) {
	JsonResponse(response, err.Code, err.Body)
}

func CustomError(response http.ResponseWriter, code int, body interface{}) {
	JsonResponse(response, code, body)
}

func NotFound(response http.ResponseWriter) {
	JsonResponse(response, http.StatusNotFound, map[string]interface{}{
		"message": "Not found",
	})
}

func Unauthorized(response http.ResponseWriter) {
	JsonResponse(response, http.StatusUnauthorized, map[string]interface{}{
		"message": "Unauthorized",
	})
}

func MethodNotAllowed(response http.ResponseWriter) {
	JsonResponse(response, http.StatusMethodNotAllowed, map[string]interface{}{
		"message": "Method not allowed",
	})
}

func BadRequest(response http.ResponseWriter) {
	JsonResponse(response, http.StatusBadRequest, map[string]interface{}{
		"message": "Bad request",
	})
}

func InternalServerError(response http.ResponseWriter) {
	JsonResponse(response, http.StatusInternalServerError, map[string]interface{}{
		"message": "Internal server error",
	})
}
