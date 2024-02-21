package userRepo

import (
	"database/sql"
	user "go-rest-template/api/user/domain"
	"go-rest-template/api/user/infra/database"
	userMapper "go-rest-template/api/user/mapper"
	"go-rest-template/internal/app/database"
)

func GetWithId(id *int) (*user.User, error) {
	result, e := database.QueryRow[user.RawUser]("SELECT * FROM "+userConstants.UserTableName+"where id=?", 1)
	if e != nil {
		return nil, e
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
