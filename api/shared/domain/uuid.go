package shared

import (
	"github.com/google/uuid"
)

type UUID struct {
	Value string
}

func (u *UUID) CreateV4() (string, error) {
	u.Value = uuid.New().String()
	return u.Value, nil
}
