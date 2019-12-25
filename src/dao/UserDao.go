package dao

import (
	"appstud.com/github-core/src/database"
	"appstud.com/github-core/src/models"
	"database/sql"
	"errors"
	"strings"
)

var idStub = "{{ID}}"
var loginStub = "{{LOGIN}}"
var passwordStub = "{{PASSWORD}}"
var selectUserStatement = "select id, login from users where id='" + idStub + "';"
var selectUserLoginStatement = "select id, login from users where login='" + loginStub + "';"
var selectAllUsersStatement = "select id, login from users;"
var selectCountUserLoginStatement = "select count(id) from users where login='" + loginStub + "';"
var selectUserLoginPasswordStatement = "select id, login from users where login='" + loginStub + "' and password='" + passwordStub + "';"
var insertUserStatement = "insert or ignore into users(id, login, password) values ('" + idStub + "','" + loginStub + "','" + passwordStub + "');"

func GetAllUsers() ([]models.DatabaseUser, error) {
	var rows, error = database.ExecuteSimpleQuery(selectAllUsersStatement)
	var result = parseUsers(rows)
	return result, error
}

func GetContainsUserLogin(login string) (bool, error) {
	var query = strings.ReplaceAll(selectUserLoginStatement, loginStub, login)
	println("##### ", query)
	var rows, error = database.ExecuteSimpleQuery(query)
	var result = make([]models.DatabaseUser, 0)

	if error != nil {
		println("##### error: ", error.Error())
		return false, error
	}
	result = parseUsers(rows)
	if len(result) == 0 {
		return false, nil
	}
	if len(result) > 1 {
		return true, errors.New("more than one user found")
	}
	return true, nil
}

func GetUserById(id string) (models.DatabaseUser, error) {
	var empty models.DatabaseUser
	var query = strings.ReplaceAll(selectUserStatement, idStub, id)
	println("##### ", query)
	var rows, error = database.ExecuteSimpleQuery(query)
	var result = make([]models.DatabaseUser, 0)
	if error != nil {
		return empty, error
	}
	result = parseUsers(rows)
	if len(result) == 0 {
		return empty, errors.New("user not found")
	}
	if len(result) > 1 {
		return empty, errors.New("more than one user found")
	}
	return result[0], nil
}

func GetUserByLoginAndPassword(login string, password string) (models.DatabaseUser, error) {
	var empty models.DatabaseUser
	var query = strings.ReplaceAll(selectUserLoginPasswordStatement, loginStub, login)
	query = strings.ReplaceAll(query, passwordStub, password)
	println("##### ", query)
	var rows, error = database.ExecuteSimpleQuery(query)
	var result = make([]models.DatabaseUser, 0)
	if error != nil {
		return empty, error
	}
	result = parseUsers(rows)
	if len(result) == 0 {
		return empty, errors.New("user not found")
	}
	if len(result) > 1 {
		return empty, errors.New("more than one user found")
	}
	return result[0], nil
}

func AddUser(id string, login string, password string) (int64, error) {
	var hasUser, cError = GetContainsUserLogin(login)
	if cError != nil {
		return 0, cError
	}
	if hasUser {
		return 0, errors.New("user laready exists")
	}

	var query = strings.ReplaceAll(insertUserStatement, idStub, id)
	query = strings.ReplaceAll(query, loginStub, login)
	query = strings.ReplaceAll(query, passwordStub, password)
	println("##### ", query)
	var result, eError = database.ExecuteSimpleUpdatableQuery(query)
	if eError != nil {
		return 0, eError
	}
	return result.RowsAffected()
}

func parseUsers(rows *sql.Rows) []models.DatabaseUser {
	var users = make([]models.DatabaseUser, 0)
	for rows.Next() {
		var id string
		var username string
		rows.Scan(&id, &username)
		users = append(users, models.DatabaseUser{
			Id:       id,
			Username: username,
		})
	}
	return users
}
