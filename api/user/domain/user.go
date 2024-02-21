package user

import (
	shared "go-rest-template/api/shared/domain"
	user "go-rest-template/api/user/domain/valueObjects"
)

type User struct {
	ID        shared.ID     `json:"id" db:"id"`
	UUID      shared.UUID   `json:"uuid" db:"pid"`
	Name      user.Name     `json:"name" db:"name"`
	Email     user.Email    `json:"email" db:"email"`
	Password  user.Password `json:"password" db:"password"`
	CreatedAt shared.Date   `json:"createdA" db:"createdAt"`
	UpdatedAt shared.Date   `json:"updatedAt" db:"updatedAt"`
}

type RawUser struct {
	ID        int    `db:"id"`
	UUID      string `db:"pid"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt string `db:"createdAt"`
	UpdatedAt string `db:"updatedAt"`
}
