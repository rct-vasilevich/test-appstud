package cache

import (
	"appstud.com/github-core/src/models"
	"time"
)

type GitHubCachedUser struct {
	User     models.GitHubUserResponse
	LastView int
}

var users = make(map[string]GitHubCachedUser)

func CheckUserInCache(user string) bool {
	if _, contains := users[user]; contains {
		return true
	}
	return false
}

func GetUserFromCache(user string) models.GitHubUserResponse {
	return users[user].User
}

func AddUserToCache(user models.GitHubUserResponse) {
	if len(feed) > 50 {
		deleteOldUser()
	}
	users[user.Login] = GitHubCachedUser{
		User:     user,
		LastView: time.Now().Second(),
	}
}

func deleteOldUser() {
	var minTime = time.Now().Second()
	var user = ""
	for index, element := range users {
		if element.LastView < minTime {
			minTime = element.LastView
			user = index
		}
	}
	delete(users, user)
}
