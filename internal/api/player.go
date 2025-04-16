package api

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/handler"
	"github.com/gin-gonic/gin"
)

func (a *API) addPlayerRoute(rg *gin.RouterGroup) {
	team := rg.Group("/player")

	team.GET("/:playerTag", handler.Player)
}
