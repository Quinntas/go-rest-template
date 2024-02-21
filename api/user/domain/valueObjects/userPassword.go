package userValueObjects

import "go-rest-template/internal/app/web"

type Password struct {
	Value string
}

func ValidatePassword(value string) (Password, *web.HttpError) {
	return Password{Value: value}, nil
}
