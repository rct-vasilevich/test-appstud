package main

import (
	"appstud.com/github-core/src/controllers"
	"appstud.com/github-core/src/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initDatabase()
	initControllers(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func initDatabase() {
	database.Init()
}

func initControllers(r *gin.Engine) {
	controllers.HelloWorldController(r)
	controllers.HealthCheckController(r)
	controllers.EasterEggsController(r)
	controllers.UserController(r)
	controllers.GitHubController(r)
}
