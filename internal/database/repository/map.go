package repository

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
)

func GetMapByName(name string) (entity.Map, error) {
	db := database.Db()
	m := entity.Map{}

	if err := db.Where("name = ?", name).First(&m).Error; err != nil {
		return m, err
	}

	return m, nil
}
