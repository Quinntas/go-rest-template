package login

import (
	userValueObjects "go-rest-template/api/user/domain/valueObjects"
	userRepo "go-rest-template/api/user/repo"
	"go-rest-template/internal/app/redis"
	"go-rest-template/internal/app/utils"
	"go-rest-template/internal/app/web"
	"net/http"
	"time"
)

func UseCase(response http.ResponseWriter, request *http.Request, decodedRequest *web.DecodedRequest[DTO]) *web.HttpError {
	email, validationErr := userValueObjects.ValidateEmail(decodedRequest.Json.Email)
	if validationErr != nil {
		return validationErr
	}

	user, err := userRepo.GetWithEmail(email.Value)
	if err != nil {
		return &web.HttpError{
			Code: http.StatusNotFound,
			Body: map[string]interface{}{
				"message": "User not found",
			},
		}
	}

	if result, err := utils.CompareEncryption(decodedRequest.Json.Password, user.Password.Value); err != nil || result == false {
		return &web.HttpError{
			Code: http.StatusUnauthorized,
			Body: map[string]interface{}{
				"message": "Invalid password",
			},
		}
	}

	privateToken, err := utils.GenerateJsonWebToken[PrivateTokenClaim](&PrivateTokenClaim{
		Id:    user.ID.Value,
		Email: user.Email.Value,
		UUID:  user.UUID.Value,
	}, TOKEN_EXPIRATION_TIME)

	err = redis.Set(TOKEN_REDIS_KEY+user.UUID.Value, privateToken, TOKEN_EXPIRATION_TIME)
	if err != nil {
		return &web.HttpError{
			Code: http.StatusInternalServerError,
			Body: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	if err != nil {
		return &web.HttpError{
			Code: http.StatusInternalServerError,
			Body: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	publicToken, err := utils.GenerateJsonWebToken[PublicTokenClaim](&PublicTokenClaim{
		UUID: user.UUID.Value,
	}, TOKEN_EXPIRATION_TIME)

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
		ExpiresIn:  int(TOKEN_EXPIRATION_TIME.Seconds()),
		ExpireDate: utils.TimeToString(time.Now().Add(TOKEN_EXPIRATION_TIME)),
	})

	return nil
}
