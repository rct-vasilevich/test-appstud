package controllers

import (
	"appstud.com/github-core/src/models"
	"github.com/gin-gonic/gin"
	"time"
)

func HealthCheckController(engine *gin.Engine) {
	engine.GET("/api/healthcheck", handleHealthCheck)
}

func handleHealthCheck(c *gin.Context) {
	c.JSON(200, models.HealthCheckResponse{
		Name:    "github-api",
		Version: "1.0",
		Time:    int64(time.Now().Second()),
	})
}
