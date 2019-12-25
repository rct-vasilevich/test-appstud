package database

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

var driver = "sqlite3"
var db *sql.DB

func Connect(path string) {
	db, _ = sql.Open(driver, path)
}

func ExecuteSimpleUpdatableQuery(query string) (sql.Result, error) {
	var result sql.Result
	if db != nil {
		return db.Exec(query)
	} else {
		return result, errors.New("database not connected")
	}
}

func ExecuteSimpleQuery(query string) (*sql.Rows, error) {
	var result *sql.Rows
	if db != nil {
		return db.Query(query)
	} else {
		return result, errors.New("database not connected")
	}
}
