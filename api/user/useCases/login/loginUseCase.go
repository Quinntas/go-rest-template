package login

import (
	"go-rest-template/internal/app/utils"
	"go-rest-template/internal/app/web"
	"net/http"
	"time"
)

func UseCase(response http.ResponseWriter, request *http.Request, decodedRequest *web.DecodedRequest[DTO]) *web.HttpError {
	_, err := utils.GenerateJsonWebToken[PrivateTokenClaim](&PrivateTokenClaim{
		Id:    5,
		Email: "caoio@gmail.com",
		UUID:  "apsokdapsokdp",
	}, time.Time{})

	if err != nil {
		return &web.HttpError{
			Code: http.StatusInternalServerError,
			Body: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	publicToken, err := utils.GenerateJsonWebToken[PublicTokenClaim](&PublicTokenClaim{
		UUID: "apsokdapsokdp",
	}, time.Time{})

	if err != nil {
		return &web.HttpError{
			Code: http.StatusInternalServerError,
			Body: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	web.JsonResponse[ResponseDTO](response, http.StatusOK, &ResponseDTO{
		Token:      publicToken,
		ExpiresIn:  3600,
		ExpireDate: "2021-01-01",
	})

	return nil
}
