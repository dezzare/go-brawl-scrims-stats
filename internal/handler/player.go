package handler

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
)

func GetAllPlayers() *[]entity.Player {
	db := database.Db()
	players := []entity.Player{}

	db.Find(&players)
	return &players
}

func GetPlayerByTag(s string) entity.Player {
	db := database.Db()
	p := entity.Player{Tag: s}

	db.First(&p)
	return p
}
