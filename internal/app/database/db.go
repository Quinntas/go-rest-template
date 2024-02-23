package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

func MustExec(query string, args ...interface{}) sql.Result {
	return db.MustExec(query, args...)
}

func QueryRow[T interface{}](query string, args ...interface{}) (*T, error) {
	var result T
	err := db.QueryRowx(query, args...).StructScan(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func Query[T interface{}](query string, args ...interface{}) ([]*T, error) {
	rows, err := db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*T, 0)
	for rows.Next() {
		var row T
		err = rows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		results = append(results, &row)
	}
	return results, nil
}

func ConnectDB(databaseUrl string) {
	db = sqlx.MustConnect("mysql", databaseUrl)
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}
