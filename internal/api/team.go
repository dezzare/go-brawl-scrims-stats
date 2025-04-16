package api

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/handler"
	"github.com/gin-gonic/gin"
)

func (a *API) addTeamRoute(rg *gin.RouterGroup) {
	team := rg.Group("/team")

	team.GET("/:name", handler.Team)
}
