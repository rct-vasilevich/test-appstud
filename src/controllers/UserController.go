package controllers

import (
	"appstud.com/github-core/src/facade"
	"appstud.com/github-core/src/models"
	"github.com/gin-gonic/gin"
)

func UserController(engine *gin.Engine) {
	engine.GET("/api/users", handleGetAllUsers)
	engine.GET("/api/users/register", handleRegisterUser)
	engine.GET("/api/users/login", handleLoginUser)
	engine.GET("/api/users/me", handleConnectedUser)
}

func handleGetAllUsers(c *gin.Context) {
	c.JSON(200, facade.GetAllUsers())
}

func handleRegisterUser(c *gin.Context) {
	var username = c.Request.URL.Query().Get("username")
	var password = c.Request.URL.Query().Get("password")
	if username == "" || password == "" {
		c.JSON(400, models.ErrorResponse{
			Message: "please check your username and password (/api/users/login?username={username}&password={password})",
		})
	} else {
		var user, error = facade.RegisterUser(username, password)
		if error == nil {
			c.JSON(200, user)
		} else {
			c.JSON(400, models.ErrorResponse{Message: error.Error()})
		}
	}
}

func handleLoginUser(c *gin.Context) {
	var username = c.Request.URL.Query().Get("username")
	var password = c.Request.URL.Query().Get("password")
	if username == "" || password == "" {
		c.JSON(400, models.ErrorResponse{
			Message: "please check your username and password (/api/users/login?username={username}&password={password})",
		})
	} else {
		var user, error = facade.LoginUser(username, password)
		if error == nil {
			c.JSON(200, user)
		} else {
			c.JSON(404, models.ErrorResponse{Message: error.Error()})
		}
	}
}

func handleConnectedUser(c *gin.Context) {
	var uuid = c.Request.URL.Query().Get("token")
	if uuid == "" {
		c.JSON(400, "please provide me a token (/api/users/me?token={token})")
	} else {
		var user, error = facade.GetConnectedUser(uuid)
		if error == nil {
			c.JSON(200, user)
		} else {
			c.JSON(404, models.ErrorResponse{Message: error.Error()})
		}
	}
}
