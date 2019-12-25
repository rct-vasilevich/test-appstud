package cache

import (
	"appstud.com/github-core/src/models"
	"time"
)

type GitHubCachedFeed struct {
	Feed     []models.GitHubEventResponse
	LastView int
}

var feed = make(map[string]GitHubCachedFeed)

func CheckFeedInCache(user string) bool {
	if _, contains := feed[user]; contains {
		return true
	}
	return false
}

func GetFeedFromCache(user string) []models.GitHubEventResponse {
	return feed[user].Feed
}

func AddFeedToCache(user string, newFeed []models.GitHubEventResponse) {
	if len(feed) > 50 {
		deleteOldFeed()
	}
	feed[user] = GitHubCachedFeed{
		Feed:     newFeed,
		LastView: time.Now().Second(),
	}
}

func deleteOldFeed() {
	var minTime = time.Now().Second()
	var user = ""
	for index, element := range feed {
		if element.LastView < minTime {
			minTime = element.LastView
			user = index
		}
	}
	delete(feed, user)
}
