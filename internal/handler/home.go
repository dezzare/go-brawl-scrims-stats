package handler

import (
	"net/http"
	"time"

	"github.com/dezzare/go-brawl-scrims-stats/internal/service"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	db := ctx.MustGet("db").(*service.DB)
	isLogged := true
	testPlayer := "#LVRRYPV"

	player, _ := db.GetPlayerByTag(testPlayer)
	battles := player.Battles
	time.Sleep(500 * time.Millisecond)

	ctx.HTML(http.StatusOK, "indexPage", gin.H{
		"playerName": player.Name,
		"loggedIn":   isLogged,
		"items":      battles,
	})

}
