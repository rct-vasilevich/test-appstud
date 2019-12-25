package models

type GitHubEventResponse struct {
	Type  string      `json:"type"`
	Actor GitHubActor `json:"actor"`
	Repo  GitHubRepo  `json:"repo"`
}

type GitHubActor struct {
	Id    int    `json:id`
	Login string `json:"login"`
}

type GitHubRepo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
