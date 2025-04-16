package repository

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
)

func (r *PostgresRepository) FindOrCreateMode(modeName string) (uint, error) {
	var mode model.Mode
	// Ignore case sensitivity
	err := r.DB.Where("LOWER(name) = LOWER(?)", modeName).First(&mode).Error
	if err == nil {
		// Found
		return mode.ID, nil
	}

	mode = model.Mode{Name: modeName}
	err = r.DB.Create(&mode).Error
	if err != nil {
		return 0, fmt.Errorf("erro ao criar modo '%s': %w", modeName, err)
	}
	return mode.ID, nil
}

// func GetModeByName(name string) (model.Mode, error) {
// 	db := database.Db()
// 	m := model.Mode{}

// 	if err := db.Where("name = ?", name).First(&m).Error; err != nil {
// 		return m, err
// 	}

// 	return m, nil
// }
