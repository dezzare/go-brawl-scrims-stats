package stats

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dezzare/go-brawl-scrims-stats/internal/client"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
	"github.com/dezzare/go-brawl-scrims-stats/internal/service"
	"github.com/dezzare/go-brawl-scrims-stats/pkg/convert"
)

func setPlayersBase(db *service.DB) error {
	fmt.Println("[SPB] Loading players to memory")
	rawPlayers, err := loadPlayersFromFile()
	if err != nil {
		return err
	}

	for _, rawPlayer := range rawPlayers.Player {
		fmt.Println("Looking in DB for player: ", rawPlayer.Tag)
		teamInDB, err := db.FindOrCreateTeam(rawPlayer.Team)
		playerInDB, err := db.GetPlayerByTag(rawPlayer.Tag)
		if err != nil && rawPlayer.Tag != "" {
			fmt.Println("Player: ", rawPlayer.Name, "not in DB")
			p := convert.RawToPlayer(&rawPlayer, teamInDB)
			err = db.CreatePlayer(p)
			if err != nil {
				fmt.Println("Error creating player from file: ", err)
			}
			continue
		}
		if playerInDB.Follow != true {
			if err = db.SetPlayerFollowStatus(playerInDB, true); err != nil {
				fmt.Println("[SPB] Error setting Follow Status: ", err)
			}
		}
		if playerInDB.TeamID != &teamInDB.ID {
			var updateData map[string]interface{}
			if rawPlayer.Team == "" {
				updateData = map[string]interface{}{"team_id": nil}
			} else {
				updateData = map[string]interface{}{"team_id": teamInDB.ID}
			}
			if err = db.UpdatePlayer(playerInDB, updateData); err != nil {
				fmt.Println("[SPB] Error setting new Team for player: ", playerInDB.Tag, teamInDB.Name, err)
			}
		}
	}
	fmt.Println("[SPB] Player base ready")
	return nil
}

func setPlayersBattlelog(db *service.DB, c *client.Client) error {
	fmt.Println("[SPB] Setting Players Battlelog")
	players, err := db.GetPlayersFollowed()
	if err != nil {
		return err
	}
	for _, p := range *players {
		if p.Follow {
			rawBl := c.GetBattleLog(p.Tag)
			_, err := convert.RawMatchesToBattlesAndSave(rawBl, db)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println("[SPB] Players Battlelog ready")
	return nil
}

func setPlayersBrawlerStat(pbs *[]model.PlayerBrawlerStat, players *[]model.Player, db *service.DB) error {
	// Create map for stack stats of each player
	var playersID []uint
	for _, player := range *players {
		playersID = append(playersID, player.ID)
	}

	var playerResults []model.PlayerResult
	if err := db.GetPlayerResults(playersID, &playerResults); err != nil {
		fmt.Println(err)
	}

	brawlers := db.GetAllBrawlers()

	playerBrawlerStats, err := convert.ToPlayerBrawlerStat(&playerResults, players, brawlers, db)
	if err != nil {
		return err
	}

	(*pbs) = playerBrawlerStats

	return nil
}

func loadPlayersFromFile() (*model.RawPlayers, error) {
	file, err := os.ReadFile("players.json")
	if err != nil {
		return nil, fmt.Errorf("Open file error: %v", err)
	}
	var result model.RawPlayers
	if err := json.Unmarshal(file, &result); err != nil {
		return nil, fmt.Errorf("Players JSON Unmarshhal error: %v", err)
	}

	return &result, nil
}

// func saveBattlelog(bl []model.Battle, db *service.DB) {
// 	for _, v := range bl {
// 		if err := db.CreateBattle(&v); err != nil {
// 			fmt.Println("Error saving battle: ", err)
// 		}
// 	}
// }
