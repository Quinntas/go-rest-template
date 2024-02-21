package createUser

import (
	shared "go-rest-template/api/shared/domain"
	user "go-rest-template/api/user/domain"
	userValueObjects "go-rest-template/api/user/domain/valueObjects"
	userRepo "go-rest-template/api/user/repo"
	"go-rest-template/internal/app/web"
	"net/http"
)

func UseCase(response http.ResponseWriter, request *http.Request, decodedRequest *web.DecodedRequest[DTO]) {
	email, err := userValueObjects.ValidateEmail(decodedRequest.Json.Email)
	if err != nil {
		web.HandleHttpError(response, *err)
	}

	name, err := userValueObjects.ValidateName(decodedRequest.Json.Name)
	if err != nil {
		web.HandleHttpError(response, *err)
	}

	password, err := userValueObjects.ValidatePassword(decodedRequest.Json.Password)
	if err != nil {
		web.HandleHttpError(response, *err)
	}

	_, dbErr := userRepo.Create(&user.User{
		Email:    email,
		Name:     name,
		Password: password,
		UUID:     shared.CreateV4(),
	})

	if dbErr != nil {
		web.JsonResponse(response, http.StatusBadRequest, map[string]interface{}{
			"message": dbErr.Error(),
		})
		return
	}

	web.JsonResponse(response, http.StatusCreated, map[string]interface{}{
		"message": "ok",
	})

	return
}
