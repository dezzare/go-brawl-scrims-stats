package repository

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
	"gorm.io/gorm/clause"
)

func (r *PostgresRepository) CreateTeam(teamToSave *model.Team) error {
	if err := r.DB.Create(&teamToSave).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) FindOrCreateTeam(teamName string) (*model.Team, error) {
	var t model.Team
	err := r.DB.Where("name = ?", teamName).First(&t).Error
	if err == nil {
		return &t, nil
	}
	t.Name = teamName
	err = r.DB.Create(&t).Error
	if err != nil {
		return nil, fmt.Errorf("erro ao criar time tag '%s': %w", t.Name, err)
	}
	return &t, nil
}

func (r *PostgresRepository) GetTeamPlayers(players *[]model.Player, teamID uint) error {
	if err := r.DB.Where("team_id = ?", teamID).Preload(clause.Associations).Find(players).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) GetTeamByName(name string) (*model.Team, error) {
	var team model.Team
	if err := r.DB.Where("name = ?", name).Find(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *PostgresRepository) GetTeamByID(teamID uint) (*model.Team, error) {
	var team model.Team
	if err := r.DB.Where("id = ?", teamID).First(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *PostgresRepository) GetTeamBattles(teamID uint) (*[]model.PlayerResult, error) {
	var playerIDs []uint
	if err := r.DB.Model(&model.Player{}).Where("team_id = ?", teamID).Pluck("id", &playerIDs).Error; err != nil {
		return nil, err
	}

	teamPlayerCount := len(playerIDs)
	var battleIDs []uint
	err := r.DB.Model(&model.PlayerResult{}).
		Select("battle_id").
		Where("player_id IN ?", playerIDs).
		Group("battle_id").
		Having("COUNT(DISTINCT player_id) = ?", teamPlayerCount).
		Pluck("battle_id", &battleIDs).Error
	if err != nil {
		return nil, err
	}

	var result []model.PlayerResult
	err = r.DB.Where("id IN ?", playerIDs).
		Preload("Player").
		Preload("Brawler").
		Preload("Battle").
		Preload("Battle.Map").
		Preload("Battle.Mode").
		Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
