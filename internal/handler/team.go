package handler

import (
	"fmt"
	"net/http"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
	"github.com/dezzare/go-brawl-scrims-stats/internal/service"
	"github.com/gin-gonic/gin"
)

func Team(ctx *gin.Context) {
	db := ctx.MustGet("db").(*service.DB)
	name := ctx.Params.ByName("name")
	team, err := db.GetTeamByName(name)
	if err != nil {
		fmt.Println(err)
	}
	var players []model.Player
	if err := db.GetTeamPlayers(&players, team.ID); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"team": name, "status": "no value"})

	}
	ctx.JSON(http.StatusOK, gin.H{"team": name, "value": players})

}
