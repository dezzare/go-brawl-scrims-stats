package repository

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateBattle recieve entity.Battle and save to DB
func CreateBattle(b entity.Battle) *entity.Battle {
	db := database.Db()

	if !haveBattle(db, &b) {
		fmt.Println("Saving battle to DB")
		if err := db.Create(&b).Error; err != nil {
			fmt.Println("Error saving battle to DB: ", err)
			return nil
		}
		fmt.Println("Success saving battle")
		return &b
	}
	return &b
}

// CreateBattleResult recieve entity.BattleResult and save to DB
func CreateBattleResult(b entity.BattleResult) {
	db := database.Db()

	if !haveBattleResult(db, &b) {
		fmt.Println("Saving BattleResult to DB")
		if err := db.Create(&b).Error; err != nil {
			fmt.Println("Error saving battleresult to DB: ", err)
		}
	}
}

func GetBattleresultByPlayerIDs(ids []uint, battleResults *[]entity.BattleResult) error {
	db := database.Db()
	if err := db.Where("player_id IN (?)", ids).Find(battleResults).Error; err != nil {
		return err
	}
	return nil
}

func GetAllBattlesByPlayerIDs(pids []uint, battles *[]entity.Battle) error {
	db := database.Db()
	if err := db.Joins("JOIN battle_results br ON br.battle_id = battles.id").
		Where("br.player_id IN (?)", pids).
		Find(battles).Error; err != nil {
		return err
	}
	return nil
}

func GetSharedBattles(players []uint, battleResults *[]entity.BattleResult) error {
	db := database.Db()
	var commonBattleIDs []uint
	// Buscar todas as batalhas em que cada jogador participou
	for k, id := range players {
		var playerBattles []uint
		err := db.Model(&entity.BattleResult{}).
			Where("player_id = ?", id).
			Pluck("battle_id", &playerBattles).Error
		if err != nil {
			return err
		}

		if k == 0 {
			// Para o primeiro jogador, inicializamos os IDs de batalha comuns
			commonBattleIDs = playerBattles
		} else {
			// Filtramos os IDs que são comuns entre os jogadores
			commonBattleIDs = intersect(commonBattleIDs, playerBattles)
		}

		// Se em algum momento a interseção for vazia, não há batalhas em comum
		if len(commonBattleIDs) == 0 {
			return nil
		}
	}

	// Buscar os BattleResult correspondentes aos BattleIDs filtrados
	err := db.Where("battle_id IN ?", commonBattleIDs).Find(&battleResults).Error
	if err != nil {
		return err
	}

	return nil
}

// haveBattle check if battle already in DB
func haveBattle(db *gorm.DB, battle *entity.Battle) bool {

	if err := db.Model(battle).Preload(clause.Associations).Where(
		"battle_time = ? AND mode = ? AND map = ?",
		battle.BattleTime, battle.Mode, battle.Map,
	).First(&battle).Error; err != nil {
		return false
	}

	return true
}

func haveBattleResult(db *gorm.DB, battle *entity.BattleResult) bool {
	var r entity.BattleResult
	result := db.Model(r).Preload(clause.Associations).Where(
		"result = ? AND brawler_id = ? AND player_id = ? AND battle_id = ?",
		battle.Result, battle.BrawlerID, battle.PlayerID, battle.BattleID,
	).First(battle)
	if result.Error != nil {
		return false
	}
	return true
}

func intersect(a, b []uint) []uint {
	set := make(map[uint]bool)
	var result []uint

	// Adiciona todos os elementos de `a` ao mapa
	for _, id := range a {
		set[id] = true
	}

	// Mantém apenas os IDs que também estão em `b`
	for _, id := range b {
		if set[id] {
			result = append(result, id)
		}
	}

	return result
}
