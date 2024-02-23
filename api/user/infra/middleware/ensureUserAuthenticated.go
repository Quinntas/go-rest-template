package userMiddleware

import (
	"go-rest-template/internal/app/utils"
	"go-rest-template/internal/app/web"
	"net/http"
	"strings"
)

func EnsureUserAuthenticated[T interface{}](next web.HandlerFunc[T]) web.HandlerFunc[T] {
	return func(response http.ResponseWriter, request *http.Request, decodedRequest *web.DecodedRequest[T]) *web.HttpError {
		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			web.Unauthorized(response)
			return nil
		}

		token := strings.Replace(authHeader, "Bearer ", "", -1)
		if token == "" {
			web.Unauthorized(response)
			return nil
		}

		utils.Print(token)

		next(response, request, decodedRequest)
		return nil
	}
}
