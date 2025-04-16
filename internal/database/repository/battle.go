package repository

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
)

// CreateBattle recieve *model.Battle and create a new record if not found
func (r *PostgresRepository) CreateBattle(battleToSave *model.Battle) error {

	if err := r.DB.Where(*battleToSave).FirstOrCreate(battleToSave).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) GetBattleByID(id string) (*model.Battle, error) {
	p := model.Battle{}
	if err := r.DB.Where("tag = ?", id).First(&p).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PostgresRepository) GetAllBattles() (*[]model.Battle, error) {
	var battles []model.Battle
	if err := r.DB.Find(&battles).Error; err != nil {
		return nil, err
	}

	return &battles, nil
}

// func haveBattleResult(db *gorm.DB, battle *model.PlayerResult) bool {
// 	var r model.PlayerResult
// 	result := db.Model(r).Preload(clause.Associations).Where(
// 		"result = ? AND brawler_id = ? AND player_id = ? AND battle_id = ?",
// 		battle.Result, battle.BrawlerID, battle.PlayerID, battle.BattleID,
// 	).First(battle)
// 	if result.Error != nil {
// 		return false
// 	}
// 	return true
// }

// func GetAllBattlesByPlayerIDs(pids []uint, battles *[]model.Battle) error {
// 	db := database.Db()
// 	if err := db.Joins("JOIN battle_results br ON br.battle_id = battles.id").
// 		Where("br.player_id IN (?)", pids).
// 		Find(battles).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
