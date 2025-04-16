package repository

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
)

func (r *PostgresRepository) FindOrCreateMap(mapName string, modeID uint) (uint, error) {
	var mapRecord model.Map
	err := r.DB.Where("LOWER(name) = LOWER(?) AND mode_id = ?", mapName, modeID).First(&mapRecord).Error
	if err == nil {
		return mapRecord.ID, nil // Encontrado
	}

	// Não encontrado, criar
	fmt.Printf("INFO: Mapa '%s' (ModoID: %d) não encontrado, criando...\n", mapName, modeID)
	var m model.Mode
	err = r.DB.Where("id = ?", modeID).First(&m).Error
	if err != nil {
		fmt.Println("Couldn't find mode: %w", err)
	}
	mapRecord = model.Map{Name: mapName, ModeID: modeID, Mode: m}
	err = r.DB.Create(&mapRecord).Error
	if err != nil {
		return 0, fmt.Errorf("erro ao criar mapa '%s' (modo %d): %w", mapName, modeID, err)
	}
	return mapRecord.ID, nil
}

// func GetMapByName(name string) (model.Map, error) {
// 	db := database.Db()
// 	m := model.Map{}

// 	if err := db.Where("name = ?", name).First(&m).Error; err != nil {
// 		return m, err
// 	}

// 	return m, nil
// }
