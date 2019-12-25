package dao

import (
	"appstud.com/github-core/src/database"
	"appstud.com/github-core/src/models"
	"errors"
	"strings"
)

var sessionIdStub = "{{SESSION_ID}}"
var userIdStub = "{{USER_ID}}"
var selectSessionStatement = "select id, user_id from sessions where id='" + sessionIdStub + "';"
var insertSessionStatement = "insert into sessions(id, user_id) values ('" + sessionIdStub + "','" + userIdStub + "');"

func GetSessionSessionId(sessionId string) (models.DatabaseSession, error) {
	var query = strings.ReplaceAll(selectSessionStatement, sessionIdStub, sessionId)
	var rows, error = database.ExecuteSimpleQuery(query)
	println("##### ", query)
	var empty models.DatabaseSession
	var result = make([]models.DatabaseSession, 0)
	if error != nil {
		println("##### " + error.Error())
		return empty, error
	}
	for rows.Next() {
		var id string
		var userId string
		rows.Scan(&id, &userId)
		result = append(result, models.DatabaseSession{
			Id:     id,
			UserId: userId,
		})
	}
	if len(result) == 0 {
		return empty, errors.New("session not found")
	}
	if len(result) > 1 {
		return empty, errors.New("more than one session found")
	}
	return result[0], nil
}

func AddSession(sessionId string, userId string) (int64, error) {
	var query = strings.ReplaceAll(insertSessionStatement, sessionIdStub, sessionId)
	query = strings.ReplaceAll(query, userIdStub, userId)
	println("##### ", query)
	var result, eError = database.ExecuteSimpleUpdatableQuery(query)
	if eError != nil {
		return 0, eError
	}
	return result.RowsAffected()
}
