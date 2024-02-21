package user

import "errors"

const (
	MaxUserNameLength = 50
	MinUserNameLength = 3
)

type Name struct {
	Value string
}

func validateName(value string) (Name, error) {
	if len(value) > MaxUserNameLength {
		return Name{}, errors.New("name is too long")
	} else if len(value) < MinUserNameLength {
		return Name{}, errors.New("name is too short")
	}
	return Name{Value: value}, nil
}
