package handler

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
)

func GetAllTeams() *[]entity.Team {
	db := database.Db()
	teams := []entity.Team{}

	db.Find(&teams)
	return &teams
}

func GetTeamByID(id uint) (entity.Team, error) {
	db := database.Db()
	t := entity.Team{}

	if err := db.Where("id = ?", id).First(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}
