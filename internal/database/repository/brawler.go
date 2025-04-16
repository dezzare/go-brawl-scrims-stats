package repository

import (
	"errors"
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
	"gorm.io/gorm"
)

// CreateBrawler recieve *model.Brawler and create a new record if not found
func (r *PostgresRepository) CreateBrawler(brawlerToSave *model.Brawler) error {

	if err := r.DB.Where(*brawlerToSave).FirstOrCreate(brawlerToSave).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) FindOrCreateBrawler(rawBrawler *model.RawBrawler) (uint, error) {
	var brawler model.Brawler

	err := r.DB.Where("ref = ?", rawBrawler.ID).First(&brawler).Error
	if err == nil {
		return brawler.ID, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("erro ao buscar brawler ref '%v': %w", rawBrawler.ID, err) // Erro inesperado
	}

	fmt.Printf("INFO: Brawler '%s' (Ref: %v) não encontrado, criando...\n", rawBrawler.Name, rawBrawler.ID)
	brawler = model.Brawler{
		Ref:  rawBrawler.ID,
		Name: rawBrawler.Name,
	}
	err = r.DB.Create(&brawler).Error
	if err != nil {
		return 0, fmt.Errorf("erro ao criar brawler ref '%d': %w", rawBrawler.ID, err)
	}
	return brawler.ID, nil
}

func (r *PostgresRepository) GetAllBrawlers() *[]model.Brawler {
	b := []model.Brawler{}
	r.DB.Find(&b)
	return &b
}

func (r *PostgresRepository) GetBrawlerNameByID(bID uint) string {
	var brawler model.Brawler
	if err := r.DB.First(&brawler, bID).Error; err != nil {
		return "Unknown"
	}
	return brawler.Name
}

// func GetBrawlerByName(name string) (model.Brawler, error) {
// 	db := database.Db()
// 	b := model.Brawler{}

// 	if err := db.Where("name = ?", name).First(&b).Error; err != nil {
// 		return b, err
// 	}

// 	return b, nil
// }

// func compareBrawlers(db *gorm.DB, battleID uint, brawlers []model.BattleResult) bool {
// 	var existingBrawlers []model.BattleResult
// 	db.Preload("Brawlers").Where("battle_id = ?", battleID).Find(&existingBrawlers)

// 	bPlayed := make(map[uint]bool)
// 	for _, b := range existingBrawlers {
// 		bPlayed[b.BrawlerID] = false
// 	}

// 	// Return true if brawlers aren't the same
// 	for _, b := range brawlers {
// 		if !bPlayed[b.BrawlerID] {
// 			return true
// 		}
// 	}

// 	return false
// }
