package repository

import (
	"appstud.com/github-core/src/models"
	"math"
)

func GetEggs() []models.HealthCheckResponse {
	return eggs
}

var eggs = []models.HealthCheckResponse{
	{
		Name: "My mom is in love with me",
		Time: -1445470140,
	},
	{
		Name: "I go to the future and my mom end up with the wrong guy",
		Time: 1445470140,
	},
	{
		Name: "I go to the past and you will not believe what happens next",
		Time: math.MinInt64,
	},
}
