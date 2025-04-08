package util

import (
	"encoding/json"
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/registry"
)

// Aux Structs to map JSON
type RawBrawler struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Power    int    `json:"power"`
	Trophies int    `json:"trophies"`
}

type RawPlayer struct {
	Tag     string     `json:"tag"`
	Name    string     `json:"name"`
	Brawler RawBrawler `json:"brawler"`
}

type RawBattle struct {
	Mode       string        `json:"mode"`
	Type       string        `json:"type"`
	Result     string        `json:"result"`
	Duration   int           `json:"duration"`
	StarPlayer *RawPlayer    `json:"starPlayer"`
	Teams      [][]RawPlayer `json:"teams"` // 2 Teams with 3 Players
}

type RawEvent struct {
	ID   int    `json:"id"`
	Mode string `json:"mode"`
	Map  string `json:"map"`
}

type RawMatch struct {
	BattleTime string    `json:"battleTime"`
	Event      RawEvent  `json:"event"`
	Battle     RawBattle `json:"battle"`
}

type RM struct {
	Items []RawMatch `json:"items"`
}

func ConvertToBattle(data []byte) []entity.Battle {
	fmt.Println("Converting battlelog []byte to struct")
	rm := parseToRawMatch(data)
	var battles []entity.Battle

	for _, v := range rm {
		b := parseRawMatchToBattle(v)
		parseRawMatchToBattleResult(v, b)
	}

	return battles
}

func parseToRawMatch(data []byte) []RawMatch {
	var rawM RM
	if err := json.Unmarshal(data, &rawM); err != nil {
		fmt.Println("Error unmarshal ToRawMatch: ", err)
	}
	return rawM.Items
}

func parseRawMatchToBattle(match RawMatch) *entity.Battle {
	battle := entity.Battle{
		Mode:       match.Event.Mode,
		Map:        match.Event.Map,
		BattleTime: match.BattleTime,
	}
	b := registry.CreateBattle(battle)
	return b
}

func parseRawMatchToBattleResult(match RawMatch, b *entity.Battle) {
	// battleResult := []entity.BattleResult{}
	for _, teamPlayers := range match.Battle.Teams {
		players := []entity.Player{}
		auxPlayer := teamPlayers[0]
		for _, rawPlayer := range teamPlayers {
			brawler, err := registry.GetBrawlerByName(rawPlayer.Brawler.Name)
			if err != nil {
				fmt.Println("Error getting brawler: ", rawPlayer.Brawler.Name)
			}
			player, err := registry.GetPlayerByTag(rawPlayer.Tag)
			if err != nil {
				rp := entity.Player{Tag: rawPlayer.Tag, Name: rawPlayer.Name}
				registry.CreatePlayer(rp)
				player, _ = registry.GetPlayerByTag(rp.Tag)
				fmt.Println("Error getting player: ", rawPlayer.Tag)
			}

			players = append(players, player)

			battleR := entity.BattleResult{
				PlayerID:  player.ID,
				BattleID:  b.ID,
				BrawlerID: brawler.ID,
				Result:    match.Battle.Result,
			}

			if !isSameTeam(auxPlayer.Tag, player.Tag) {
				battleR.Result = changeResult(match.Battle.Result)
			}

			registry.CreateBattleResult(battleR)
		}
	}

	// battle.Results = battleResult
	// return battle
}

func changeResult(r string) string {
	switch r {
	case "victory":
		return "defeat"
	case "defeat":
		return "victory"
	default:
		return "draw"
	}
}
