package userValueObjects

import "go-rest-template/internal/app/web"

type Email struct {
	Value string
}

func ValidateEmail(value string) (Email, *web.HttpError) {
	return Email{Value: value}, nil
}
