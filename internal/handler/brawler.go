package handler

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"gorm.io/gorm"
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

func compareBrawlers(db *gorm.DB, battleID uint, brawlers []entity.BattleResult) bool {
	var existingBrawlers []entity.BattleResult
	db.Preload("Brawlers").Where("battle_id = ?", battleID).Find(&existingBrawlers)

	bPlayed := make(map[uint]bool)
	for _, b := range existingBrawlers {
		bPlayed[b.BrawlerID] = true
	}

	// Return false if brawlers aren't the same
	for _, b := range brawlers {
		if !bPlayed[b.BrawlerID] {
			return false
		}
	}

	return true
}

func GetBrawlerNameByID(bid uint) string {
	db := database.Db()
	var brawler entity.Brawler
	if err := db.First(&brawler, bid).Error; err != nil {
		return "Unknown"
	}
	return brawler.Name
}
