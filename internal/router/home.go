package router

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/handler"
	"github.com/gin-gonic/gin"
)

func addHomeRoute(rg *gin.RouterGroup) {
	home := rg.Group("/")

	home.GET("/", handler.Home)
}
