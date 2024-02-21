package userValueObjects

import (
	"go-rest-template/internal/app/web"
)

const (
	MaxUserNameLength = 50
	MinUserNameLength = 3
)

type Name struct {
	Value string
}

func ValidateName(value string) (Name, *web.HttpError) {
	nameError := web.HttpError{
		Code: 422,
		Body: map[string]interface{}{
			"message": "Invalid name",
			"key":     "name",
		},
	}
	if len(value) > MaxUserNameLength {
		return Name{}, &nameError
	} else if len(value) < MinUserNameLength {
		return Name{}, &nameError
	}
	return Name{Value: value}, nil
}
