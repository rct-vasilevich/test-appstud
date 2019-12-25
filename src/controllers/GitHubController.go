package controllers

import (
	"appstud.com/github-core/src/facade"
	"github.com/gin-gonic/gin"
)

func GitHubController(engine *gin.Engine) {
	engine.GET("/api/github/feed", handleGitHubFeed)
	engine.GET("/api/github/user/:actorlogin", handleGitHubUser)
}

func handleGitHubFeed(c *gin.Context) {
	var username = c.Request.URL.Query().Get("username")
	if username == "" {
		c.JSON(400, "please provide me a valid github username (/api/github/feed?username={username})")
	} else {
		var feed, error = facade.GetGitHubUserFeed(username)
		if error != nil {
			c.JSON(400, error.Error())
		} else {
			c.JSON(200, feed)
		}
	}
}

func handleGitHubUser(c *gin.Context) {
	var username = c.Param("actorlogin")
	var user, error = facade.GetGitHubUser(username)
	if error != nil {
		c.JSON(400, error.Error())
	} else {
		c.JSON(200, user)
	}
}
