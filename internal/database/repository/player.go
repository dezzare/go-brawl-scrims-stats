package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
)

func GetAllPlayers() *[]entity.Player {
	db := database.Db()
	players := []entity.Player{}

	db.Find(&players)
	return &players
}

func GetPlayerByTag(s string) (entity.Player, error) {
	db := database.Db()
	p := entity.Player{}
	if err := db.Where("tag = ?", s).First(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

func CreatePlayer(p entity.Player) {
	db := database.Db()
	fmt.Println("Creating in DB: ", p.ToString())
	if err := db.Create(&p).Error; err != nil {
		fmt.Println("Error creating player in DB: ", err)
	}
}

func SavePlayer(p entity.Player) {
	db := database.Db()
	fmt.Println("Saving in db: ", p.ToString())
	if err := db.Save(&p).Error; err != nil {
		fmt.Println("Error saving player in DB: ", err)
	}

}

func GetPlayersFromFile() (*entity.Players, error) {
	file, err := os.ReadFile("players.json")
	if err != nil {
		return nil, fmt.Errorf("Open file error: %v", err)
	}
	var result entity.Players
	if err := json.Unmarshal(file, &result); err != nil {
		return nil, fmt.Errorf("Players JSON Unmarshhal error: %v", err)
	}

	return &result, nil
}

func GetTeamVictories[T interface{}](t *T) error {
	db := database.Db()
	err := db.Table("battle_results").
		Select("players.team, battle_results.battle_id, battles.*, battle_results.result, players.id as player_id, players.name as player_name, brawlers.id AS brawler_id, brawlers.name as brawler_name").
		Joins("JOIN players ON players.id = battle_results.player_id").
		Joins("JOIN battles ON battles.id = battle_results.battle_id").
		Joins("JOIN brawlers ON brawlers.id = battle_results.brawler_id").
		Where("battle_results.result = ?", "victory").
		Scan(&t).Error

	if err != nil {
		return err
	}
	return nil
}

func GetTeamPlayers(players *[]entity.Player, teamName string) error {
	db := database.Db()
	if err := db.Where("team = ?", teamName).Find(players).Error; err != nil {
		return err
	}
	return nil
}
