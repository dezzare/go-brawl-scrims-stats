package handler

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
)

func GetAllBrawlers() []entity.Brawler {
	db := database.Db()
	b := []entity.Brawler{}

	db.Find(&b)
	return b
}

func GetBrawlerByName(name string) (entity.Brawler, error) {
	db := database.Db()
	b := entity.Brawler{}

	if err := db.Where("name = ?", name).First(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}
