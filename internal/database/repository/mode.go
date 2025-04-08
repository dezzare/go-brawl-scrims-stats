package repository

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
)

func GetModeByName(name string) (entity.Mode, error) {
	db := database.Db()
	m := entity.Mode{}

	if err := db.Where("name = ?", name).First(&m).Error; err != nil {
		return m, err
	}

	return m, nil
}
