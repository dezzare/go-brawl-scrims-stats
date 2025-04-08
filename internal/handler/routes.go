package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/repository"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	isLogged := true
	testPlayer := "#LVRRYPV"

	player, _ := repository.GetPlayerByTag(testPlayer)
	battles := player.Battles
	time.Sleep(500 * time.Millisecond)

	c.HTML(http.StatusOK, "indexPage", gin.H{
		"playerName": player.Name,
		"loggedIn":   isLogged,
		"items":      battles,
	})

}
func Player(c *gin.Context) {
	tag := c.Params.ByName("playerTag")
	tag = strings.ToUpper(tag)
	player, err := repository.GetPlayerByTag(tag)
	if err != nil {
		c.JSON(http.StatusNotFound, player)
	}
	c.JSON(http.StatusOK, player)

}
func Team(c *gin.Context) {
	name := c.Params.ByName("name")
	var players []entity.Player
	if err := repository.GetTeamPlayers(&players, name); err != nil {
		c.JSON(http.StatusOK, gin.H{"team": name, "status": "no value"})

	}
	c.JSON(http.StatusOK, gin.H{"team": name, "value": players})

}
