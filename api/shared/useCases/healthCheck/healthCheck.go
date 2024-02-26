package healthCheck

import (
	"go-rest-template/internal/app/web"
	"net/http"
)

func UseCase(response http.ResponseWriter, request *http.Request, decodedRequest *web.DecodedRequest[any]) *web.HttpError {
	web.JsonResponse(response, http.StatusOK, &map[string]interface{}{
		"message": "ok",
	})
	return nil
}
