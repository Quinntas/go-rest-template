package userRepo

import (
	"database/sql"
	user "go-rest-template/api/user/domain"
	"go-rest-template/api/user/infra/database"
	userMapper "go-rest-template/api/user/mapper"
	"go-rest-template/internal/app/database"
)

func GetWithEmail(email string) (*user.User, error) {
	result, err := database.QueryRow[user.RawUser]("SELECT * FROM "+userConstants.UserTableName+" WHERE email = ?;", email)
	if err != nil {
		return nil, err
	}
	return userMapper.ToDomain(result), nil
}

func Create(user *user.User) (sql.Result, error) {
	return database.Exec("INSERT INTO "+userConstants.UserTableName+" (pid, name, email, password) VALUES (?, ?, ?, ?)",
		user.UUID.Value,
		user.Name.Value,
		user.Email.Value,
		user.Password.Value)
}
