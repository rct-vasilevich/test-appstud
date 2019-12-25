package models

type LoggedInUserResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
