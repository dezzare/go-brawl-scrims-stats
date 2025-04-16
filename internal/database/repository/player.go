package repository

import (
	"errors"
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *PostgresRepository) CreatePlayer(playerToSave *model.Player) error {

	if err := r.DB.Where(*playerToSave).FirstOrCreate(playerToSave).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) FindOrCreatePlayer(rawPlayer *model.RawPlayer) (uint, error) {
	var player model.Player
	err := r.DB.Where("tag = ?", rawPlayer.Tag).First(&player).Error
	if err == nil {
		return player.ID, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("erro ao buscar jogador tag '%s': %w", rawPlayer.Tag, err) // Erro inesperado
	}

	player = model.Player{
		Tag:  rawPlayer.Tag,
		Name: rawPlayer.Name,
	}
	err = r.DB.Create(&player).Error
	if err != nil {
		return 0, fmt.Errorf("erro ao criar jogador tag '%s': %w", rawPlayer.Tag, err)
	}
	return player.ID, nil
}

func (r *PostgresRepository) GetPlayerByTag(s string) (*model.Player, error) {
	p := model.Player{}
	if err := r.DB.Where("tag = ?", s).Preload(clause.Associations).First(&p).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PostgresRepository) GetPlayersFollowed() (*[]model.Player, error) {
	p := []model.Player{}
	if err := r.DB.Where("follow = ?", true).Preload(clause.Associations).Find(&p).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PostgresRepository) GetAllPlayers() (*[]model.Player, error) {
	players := []model.Player{}
	if err := r.DB.Find(&players).Error; err != nil {
		return nil, err

	}

	return &players, nil
}

func (r *PostgresRepository) SetPlayerFollowStatus(p *model.Player, followStatus bool) error {
	var player model.Player
	if err := r.DB.Model(&player).Where("id = ?", p.ID).Update("follow", followStatus).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) GetPlayersResults(playersIDs []uint, playerResults *[]model.PlayerResult) error {
	if err := r.DB.Where("player_id IN (?)", playersIDs).Find(playerResults).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) UpdatePlayer(p *model.Player, updateData map[string]interface{}) error {
	if err := r.DB.Model(&model.Player{}).Where("id = ?", p.ID).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

// func SavePlayer(p model.Player) {
// 	db := database.Db()
// 	fmt.Println("Saving in db: ", p.ToString())
// 	if err := db.Save(&p).Error; err != nil {
// 		fmt.Println("Error saving player in DB: ", err)
// 	}

// }

// func GetTeamVictories[T interface{}](t *T) error {
// 	db := database.Db()
// 	err := db.Table("battle_results").
// 		Select("players.team, battle_results.battle_id, battles.*, battle_results.result, players.id as player_id, players.name as player_name, brawlers.id AS brawler_id, brawlers.name as brawler_name").
// 		Joins("JOIN players ON players.id = battle_results.player_id").
// 		Joins("JOIN battles ON battles.id = battle_results.battle_id").
// 		Joins("JOIN brawlers ON brawlers.id = battle_results.brawler_id").
// 		Where("battle_results.result = ?", "victory").
// 		Scan(&t).Error

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
