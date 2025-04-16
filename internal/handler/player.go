package handler

import (
	"net/http"
	"strings"

	"github.com/dezzare/go-brawl-scrims-stats/internal/service"
	"github.com/gin-gonic/gin"
)

func Player(ctx *gin.Context) {
	db := ctx.MustGet("db").(*service.DB)
	tag := ctx.Params.ByName("playerTag")
	tag = strings.ToUpper(tag)
	player, err := db.GetPlayerByTag(tag)
	if err != nil {
		ctx.JSON(http.StatusNotFound, player)
	}
	ctx.JSON(http.StatusOK, player)
}
