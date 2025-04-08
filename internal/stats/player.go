package stats

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/client"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/repository"
	"github.com/dezzare/go-brawl-scrims-stats/pkg/util"
)

type TeamStat struct {
	Team          string
	BattleResults []entity.BattleResult
	Players       []PlayerBrawlerStat
	Brawlers      []BrawlerStat
}

type PlayerBrawlerStat struct {
	PlayerName string
	Brawlers   []BrawlerStat
}

type BrawlerStat struct {
	Name      string
	Victories uint
	Defeat    uint
	Draw      uint
}

func loadPlayers() {
	fmt.Println("Loading players to memory")
	players, err := repository.GetPlayersFromFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range players.Player {
		fmt.Println("Looking in DB for player: ", v.Tag)
		p, err := repository.GetPlayerByTag(v.Tag)
		v.Follow = true
		if err != nil && v.Tag != "" {
			fmt.Println("Player: ", v, "not in DB")
			c := client.ClientConn()
			r := util.ConvertToPlayer(c.GetPlayer(v.Tag))
			r.Update(v)
			repository.CreatePlayer(r)
			continue
		}
		p.Update(v)
		repository.SavePlayer(p)
	}
}

func requestPlayerBattlelog(playerTag string) []entity.Battle {
	c := client.ClientConn()
	r := util.ConvertToBattle(c.GetBattleLog(playerTag))
	return r
}

func setPlayersBattlelog() {
	fmt.Println("Setting players DB")
	loadPlayers()
	players := repository.GetAllPlayers()

	for _, v := range *players {
		if v.Follow {
			bl := requestPlayerBattlelog(v.Tag)
			saveBattlelog(bl)
		}
	}
	return
}

func saveBattlelog(bl []entity.Battle) {
	for _, v := range bl {
		repository.CreateBattle(v)
	}
}

func setPlayersBrawlerStat(pbs *[]PlayerBrawlerStat, players *[]entity.Player) error {
	// Create map for stack stats of each player
	var playersID []uint
	for _, player := range *players {
		playersID = append(playersID, player.ID)
	}

	var battleResults []entity.BattleResult
	if err := repository.GetBattleresultByPlayerIDs(playersID, &battleResults); err != nil {
		fmt.Println(err)
	}

	brawlers := repository.GetAllBrawlers()

	playerBrawlerStats, err := convertToPlayerBrawlerStat(&battleResults, players, &brawlers)
	if err != nil {
		return err
	}

	(*pbs) = playerBrawlerStats

	return nil

}

func convertToPlayerBrawlerStat(battleResults *[]entity.BattleResult, players *[]entity.Player, brawlers *[]entity.Brawler) ([]PlayerBrawlerStat, error) {
	// Create map for stack stats of each player
	playerMap := make(map[uint]entity.Player)
	for _, player := range *players {
		playerMap[player.ID] = player
	}

	pbm := make(map[string]*PlayerBrawlerStat)
	// Take all battleResult to PlayerBrawlerStat
	for _, v := range *battleResults {
		player, exists := playerMap[v.PlayerID]
		if !exists {
			continue
		}

		brawlerName := getBrawlerNameByID(brawlers, v.BrawlerID)
		// Add player to map if not in
		if _, exists := pbm[player.Name]; !exists {
			pbm[player.Name] = &PlayerBrawlerStat{
				PlayerName: player.Name,
				Brawlers:   []BrawlerStat{},
			}
		}

		// Check if brawler in player list
		var brawlerStat *BrawlerStat
		for k := range pbm[player.Name].Brawlers {
			if pbm[player.Name].Brawlers[k].Name == brawlerName {
				brawlerStat = &pbm[player.Name].Brawlers[k]
				break
			}
		}

		// Add brawler if not in
		if brawlerStat == nil {
			newBrawler := BrawlerStat{Name: brawlerName}
			pbm[player.Name].Brawlers = append(pbm[player.Name].Brawlers, newBrawler)
			brawlerStat = &pbm[player.Name].Brawlers[len(pbm[player.Name].Brawlers)-1]
		}

		// Update stats
		switch v.Result {
		case "victory":
			brawlerStat.Victories++
		case "defeat":
			brawlerStat.Defeat++
		case "draw":
			brawlerStat.Draw++
		}
	}
	var playerStat []PlayerBrawlerStat
	// Convert map to slice and add to PlayerStat
	for _, bs := range pbm {
		playerStat = append(playerStat, *bs)
	}

	return playerStat, nil
}

func getBrawlerNameByID(brawlers *[]entity.Brawler, id uint) string {
	for _, b := range *brawlers {
		if b.ID == id {
			return b.Name
		}
	}
	return ""
}

func getBrawlerStat(pbs *[]PlayerBrawlerStat, bs *[]BrawlerStat) error {
	pbm := make(map[string]*BrawlerStat)
	for _, player := range *pbs {
		for _, brawler := range player.Brawlers {
			if _, exists := pbm[brawler.Name]; !exists {
				pbm[brawler.Name] = &BrawlerStat{
					Name: brawler.Name,
				}
			}
			pbm[brawler.Name].Victories = pbm[brawler.Name].Victories + brawler.Victories
			pbm[brawler.Name].Draw = pbm[brawler.Name].Draw + brawler.Draw
			pbm[brawler.Name].Defeat = pbm[brawler.Name].Defeat + brawler.Defeat
		}
	}

	var brawlerStat []BrawlerStat
	for _, v := range pbm {
		brawlerStat = append(brawlerStat, *v)
	}
	(*bs) = brawlerStat

	return nil
}
