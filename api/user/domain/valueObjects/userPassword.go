package userValueObjects

import (
	"go-rest-template/internal/app/utils"
	"go-rest-template/internal/app/web"
	"net/http"
)

type Password struct {
	Value string
}

func (p *Password) Encrypt() *Password {
	p.Value, _ = utils.GenerateDefaultEncryption(p.Value)
	return p
}

func ValidatePassword(value string) (*Password, *web.HttpError) {
	if len(value) < 6 {
		return nil, &web.HttpError{
			Code: http.StatusUnprocessableEntity,
			Body: map[string]interface{}{
				"message": "Password is too short",
				"key":     "password",
			},
		}
	} else if len(value) > 20 {
		return nil, &web.HttpError{
			Code: http.StatusUnprocessableEntity,
			Body: map[string]interface{}{
				"message": "Password is too long",
				"key":     "password",
			},
		}
	}
	return &Password{Value: value}, nil
}
