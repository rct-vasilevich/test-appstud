package controllers

import (
	"appstud.com/github-core/src/repository"
	"github.com/gin-gonic/gin"
	"strconv"
)

func EasterEggsController(engine *gin.Engine) {
	engine.GET("/api/timemachine/logs/mcfly", handleEasterEggs)
}

func handleEasterEggs(c *gin.Context) {
	var eggs = repository.GetEggs()
	for index := range eggs {
		version := float64(index + 1)
		eggs[index].Version = strconv.FormatFloat(version, 'f', 1, 64)
	}

	c.JSON(200, eggs)
}
