package facade

import (
	"appstud.com/github-core/src/cache"
	"appstud.com/github-core/src/models"
	"appstud.com/github-core/src/models/githubmodel"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{Timeout: 1 * time.Minute}

var githubUserTemplate = "https://api.github.com/users/" + templateUsernameStub
var githubFeedTemplate = "https://api.github.com/users/" + templateUsernameStub + "/events"
var githubPublicReposTemplate = "https://api.github.com/users/" + templateUsernameStub + "/repos"
var githubPublicGistsTemplate = "https://api.github.com/users/" + templateUsernameStub + "/gists"
var githubFollowersTemplate = "https://api.github.com/users/" + templateUsernameStub + "/followers"
var githubFollowingTemplate = "https://api.github.com/users/" + templateUsernameStub + "/following"
var templateUsernameStub = "{{username}}"

func GetGitHubUser(username string) (models.GitHubUserResponse, error) {
	if cache.CheckUserInCache(username) {
		return cache.GetUserFromCache(username), nil
	}

	var gitHubUser models.GitHubUserResponse

	var userUrl = strings.ReplaceAll(githubUserTemplate, templateUsernameStub, username)
	var gitHubApiUser githubmodel.GitHubApiUser
	var userError = getResponseObject(userUrl, &gitHubApiUser)

	var repos, reposError = getGitHubUserRepos(username)
	var gists, gistsError = getGitHubUserGists(username)
	var followers, followersError = getGitHubUserFollowers(username)
	var following, followingError = getGitHubUserFollowing(username)

	if userError != nil || reposError != nil || gistsError != nil || followersError != nil || followingError != nil {
		return gitHubUser, errors.New("error in parsing user details")
	}

	var details = models.UserDetails{
		Repos:     repos,
		Gists:     gists,
		Followers: followers,
		Following: following,
	}

	gitHubUser = models.GitHubUserResponse{
		Id:      gitHubApiUser.ID,
		Login:   gitHubApiUser.Login,
		Avatar:  gitHubApiUser.AvatarURL,
		Details: details,
	}
	cache.AddUserToCache(gitHubUser)
	return gitHubUser, nil
}

func GetGitHubUserFeed(username string) ([]models.GitHubEventResponse, error) {
	if cache.CheckFeedInCache(username) {
		return cache.GetFeedFromCache(username), nil
	}

	var events []models.GitHubEventResponse

	var eventsUrl = strings.ReplaceAll(githubFeedTemplate, templateUsernameStub, username)
	var gitHubApiEvents []githubmodel.GitHubApiEvent
	var eventsError = getResponseObject(eventsUrl, &gitHubApiEvents)

	if eventsError != nil {
		return events, eventsError
	}

	events = make([]models.GitHubEventResponse, 0)

	for _, element := range gitHubApiEvents {
		var actor = models.GitHubActor{
			Id:    element.Actor.ID,
			Login: element.Actor.Login,
		}

		var repo = models.GitHubRepo{
			Id:   element.Repo.ID,
			Name: element.Repo.Name,
		}
		events = append(events, models.GitHubEventResponse{
			Type:  element.Type,
			Actor: actor,
			Repo:  repo,
		})
	}
	cache.AddFeedToCache(username, events)
	return events, nil
}

func getGitHubUserRepos(username string) ([]models.PublicRepo, error) {
	var repos = make([]models.PublicRepo, 0)
	var reposUrl = strings.ReplaceAll(githubPublicReposTemplate, templateUsernameStub, username)
	var gitHubApiRepos []githubmodel.GitHubApiRepo
	var userReposError = getResponseObject(reposUrl, &gitHubApiRepos)

	if userReposError != nil {
		return repos, errors.New("error in parsing user details")
	}

	for _, element := range gitHubApiRepos {
		repos = append(repos, models.PublicRepo{
			Name: element.Name,
			Url:  element.URL,
		})
	}

	return nil, nil
}

func getGitHubUserGists(username string) ([]models.PublicGist, error) {
	var gists = make([]models.PublicGist, 0)
	var gistsUrl = strings.ReplaceAll(githubPublicGistsTemplate, templateUsernameStub, username)
	var gitHubApiGists []githubmodel.GitHubApiGist
	var userGistsError = getResponseObject(gistsUrl, &gitHubApiGists)

	if userGistsError != nil {
		return gists, errors.New("error in parsing user details")
	}

	for _, element := range gitHubApiGists {
		gists = append(gists, models.PublicGist{
			Id:  element.ID,
			Url: element.URL,
		})
	}

	return gists, nil
}

func getGitHubUserFollowers(username string) ([]models.GitHubSimpleUserDetails, error) {
	var followers = make([]models.GitHubSimpleUserDetails, 0)
	var followersUrl = strings.ReplaceAll(githubFollowersTemplate, templateUsernameStub, username)
	var gitHubApiFollowers []githubmodel.GitHubApiSimpleUser
	var userFollowersError = getResponseObject(followersUrl, &gitHubApiFollowers)

	if userFollowersError != nil {
		return followers, errors.New("error in parsing user details")
	}

	for _, element := range gitHubApiFollowers {
		followers = append(followers, models.GitHubSimpleUserDetails{
			Login: element.Login,
			Url:   element.URL,
		})
	}

	return followers, nil
}

func getGitHubUserFollowing(username string) ([]models.GitHubSimpleUserDetails, error) {
	var following = make([]models.GitHubSimpleUserDetails, 0)
	var followingUrl = strings.ReplaceAll(githubFollowingTemplate, templateUsernameStub, username)
	var gitHubApiFollowing []githubmodel.GitHubApiSimpleUser
	var userFollowingError = getResponseObject(followingUrl, &gitHubApiFollowing)

	if userFollowingError != nil {
		return following, errors.New("error in parsing user details")
	}

	for _, element := range gitHubApiFollowing {
		following = append(following, models.GitHubSimpleUserDetails{
			Login: element.Login,
			Url:   element.URL,
		})
	}

	return following, nil
}

func getResponseObject(url string, responseObject interface{}) error {
	var response, error = client.Get(url)
	if error != nil {
		return error
	}
	defer response.Body.Close()
	var body, readError = ioutil.ReadAll(response.Body)
	if readError != nil {
		return readError
	}
	return json.Unmarshal(body, responseObject)
}
