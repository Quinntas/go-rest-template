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

func ValidateName(value string) (*Name, *web.HttpError) {
	if len(value) > MaxUserNameLength {
		return nil, &web.HttpError{
			Code: 422,
			Body: map[string]interface{}{
				"message": "Name is too long",
				"key":     "name",
			},
		}
	} else if len(value) < MinUserNameLength {
		return nil, &web.HttpError{
			Code: 422,
			Body: map[string]interface{}{
				"message": "Name is too short",
				"key":     "name",
			},
		}
	}
	return &Name{Value: value}, nil
}
