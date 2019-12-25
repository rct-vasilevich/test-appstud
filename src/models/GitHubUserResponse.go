package models

type GitHubUserResponse struct {
	Id      int         `json:"id"`
	Login   string      `json:"login"`
	Avatar  string      `json:"avatar"`
	Details UserDetails `json:"details"`
}

type UserDetails struct {
	Repos     []PublicRepo              `json:"public_repos"`
	Gists     []PublicGist              `json:"public_gists"`
	Followers []GitHubSimpleUserDetails `json:"followers"`
	Following []GitHubSimpleUserDetails `json:"following"`
}

type PublicRepo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PublicGist struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type GitHubSimpleUserDetails struct {
	Login string `json:"login"`
	Url   string `json:"url"`
}
