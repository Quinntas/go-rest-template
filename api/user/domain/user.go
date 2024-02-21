package user

import (
	shared "go-rest-template/api/shared/domain"
	user "go-rest-template/api/user/domain/valueObjects"
)

type User struct {
	ID        shared.ID     `json:"id" db:"id"`
	UUID      shared.UUID   `json:"uuid" db:"uuid"`
	Name      user.Name     `json:"name" db:"name"`
	Email     user.Email    `json:"email" db:"email"`
	Password  user.Password `json:"password" db:"password"`
	CreatedAt shared.Date   `json:"createdA" db:"createdAt"`
	UpdatedAt shared.Date   `json:"updatedAt" db:"updatedAt"`
}
