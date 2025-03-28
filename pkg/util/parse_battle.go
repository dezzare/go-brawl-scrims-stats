package util

import (
	"encoding/json"
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"github.com/dezzare/go-brawl-scrims-stats/internal/handler"
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

func ConvertToBattle(data []byte) []entity.Battle {
	rm := parseToRawMatch(data)
	var battles []entity.Battle

	for _, v := range rm {
		battles = append(battles, parseToBattle(v))
	}

	return battles
}

func parseToRawMatch(data []byte) []RawMatch {
	var rawM []RawMatch
	if err := json.Unmarshal(data, &rawM); err != nil {
		fmt.Println(err)
	}
	return rawM
}

func parseToBattle(match RawMatch) entity.Battle {

	mode, _ := handler.GetModeByName(match.Event.Mode)
	mapS, _ := handler.GetMapByName(match.Event.Map)
	battle := entity.Battle{
		Mode:       mode,
		Map:        mapS,
		Result:     match.Battle.Result,
		BattleTime: match.BattleTime,
	}

	brawlersPlayed := []entity.BrawlerPlayed{}

	for _, teamPlayers := range match.Battle.Teams {
		players := []entity.Player{}

		for _, rawPlayer := range teamPlayers {
			brawler, _ := handler.GetBrawlerByName(rawPlayer.Brawler.Name)
			player := handler.GetPlayerByTag(rawPlayer.Tag)

			players = append(players, player)

			brawlerPlayed := entity.BrawlerPlayed{
				Player:    player,
				BrawlerID: brawler.ID,
				BattleID:  battle.ID,
			}
			brawlersPlayed = append(brawlersPlayed, brawlerPlayed)
		}

	}

	battle.Brawlers = brawlersPlayed
	return battle
}
