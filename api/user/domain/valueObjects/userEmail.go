package userValueObjects

import (
	"go-rest-template/internal/app/web"
	"net/http"
	"regexp"
)

type Email struct {
	Value string
}

func ValidateEmail(value string) (*Email, *web.HttpError) {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if result := emailRegex.MatchString(value); result == false {
		return nil, &web.HttpError{
			Code: http.StatusUnprocessableEntity,
			Body: map[string]interface{}{
				"message": "Invalid email",
				"key":     "email",
			},
		}
	}
	return &Email{Value: value}, nil
}
