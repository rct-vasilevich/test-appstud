package facade

import (
	"appstud.com/github-core/src/dao"
	"appstud.com/github-core/src/models"
	"appstud.com/github-core/src/util"
	"errors"
)

func RegisterUser(login string, password string) (models.LoggedInUserResponse, error) {
	var exists, error = dao.GetContainsUserLogin(login)
	var user models.LoggedInUserResponse
	if error != nil {
		return user, error
	}
	if exists {
		return user, errors.New("user already exists")
	}
	var userId = util.GenerateUuid()
	var i, err = dao.AddUser(userId, login, password)
	if err != nil {
		return user, err
	}
	if i != 1 {
		return user, errors.New(string(i) + " rows were affected")
	}
	var sessionId = util.GenerateUuid()
	dao.AddSession(sessionId, userId)
	return models.LoggedInUserResponse{
		Username: login,
		Token:    sessionId,
	}, nil
}

func LoginUser(login string, password string) (models.LoggedInUserResponse, error) {
	var databaseUser, error = dao.GetUserByLoginAndPassword(login, password)
	var user models.LoggedInUserResponse
	if error != nil {
		return user, error
	}
	var sessionId = util.GenerateUuid()
	dao.AddSession(sessionId, databaseUser.Id)
	user = models.LoggedInUserResponse{
		Username: databaseUser.Username,
		Token:    sessionId,
	}
	return user, nil
}

func GetAllUsers() []models.SimpleUserResponse {
	var databaseUsers, error = dao.GetAllUsers()
	users := make([]models.SimpleUserResponse, 0)
	if error == nil {
		for _, element := range databaseUsers {
			users = append(users, models.SimpleUserResponse{Username: element.Username})
		}
	}
	return users
}

func GetConnectedUser(uuid string) (models.SimpleUserResponse, error) {
	var session, sessionError = dao.GetSessionSessionId(uuid)
	var user models.SimpleUserResponse
	if sessionError != nil {
		return user, sessionError
	}
	var databaseUser, userError = dao.GetUserById(session.UserId)
	if userError != nil {
		return user, userError
	}
	user = models.SimpleUserResponse{Username: databaseUser.Username}
	return user, nil
}
