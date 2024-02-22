package createUser

import (
	shared "go-rest-template/api/shared/domain"
	user "go-rest-template/api/user/domain"
	userValueObjects "go-rest-template/api/user/domain/valueObjects"
	userRepo "go-rest-template/api/user/repo"
	"go-rest-template/internal/app/web"
	"net/http"
)

func UseCase(response http.ResponseWriter, request *http.Request, decodedRequest *web.DecodedRequest[DTO]) *web.HttpError {
	email, err := userValueObjects.ValidateEmail(decodedRequest.Json.Email)
	if err != nil {
		return err
	}

	name, err := userValueObjects.ValidateName(decodedRequest.Json.Name)
	if err != nil {
		return err
	}

	password, err := userValueObjects.ValidatePassword(decodedRequest.Json.Password)
	if err != nil {
		return err
	}

	_, dbErr := userRepo.Create(&user.User{
		Email:    *email,
		Name:     *name,
		Password: *password.Encrypt(),
		UUID:     shared.CreateV4(),
	})

	if dbErr != nil {
		return &web.HttpError{
			Code: http.StatusBadRequest,
			Body: map[string]interface{}{
				"message": dbErr.Error(),
			},
		}
	}

	web.JsonResponse(response, http.StatusCreated, &map[string]interface{}{
		"message": "ok",
	})

	return nil
}
