package shared

import (
	"github.com/google/uuid"
)

type UUID struct {
	Value string
}

func CreateV4() UUID {
	return UUID{Value: uuid.New().String()}
}
