package stats

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
	"github.com/dezzare/go-brawl-scrims-stats/internal/service"
)

func showTeamStats(teamID uint, db *service.DB) {

	teamStat, err := getTeamStat(teamID, db)
	if err != nil {
		fmt.Println("Error getting PlayerStats")
	}

	h1 := fmt.Sprintf("\n🏆 Team: %s\n", teamStat.Team)
	h1 = h1 + fmt.Sprintln("────────────────────────────────────")
	h2 := fmt.Sprintln("👥 Players & Brawlers:")
	h3 := ""
	for _, player := range teamStat.Players {
		p := fmt.Sprintf(" 🔹 Player: %s\n", player.PlayerName)
		p = p + getBrawlerStatsStr(player.Brawlers)
		h3 = h3 + p
	}
	h4 := ""
	b := fmt.Sprintln("Team Brawlers:")
	b = b + getBrawlerStatsStr(teamStat.Brawlers)
	h4 = h4 + b

	fmt.Printf("%v\n%v %v\n%v", h1, h2, h3, h4)
	fmt.Println("────────────────────────────────────")

}

func getTeamStat(teamID uint, db *service.DB) (*model.TeamStat, error) {
	var players []model.Player
	if err := db.GetTeamPlayers(&players, teamID); err != nil {
		return nil, err
	}
	var playersID []uint
	for _, p := range players {
		playersID = append(playersID, p.ID)
	}

	var pbs []model.PlayerBrawlerStat
	if err := setPlayersBrawlerStat(&pbs, &players, db); err != nil {
		return nil, err
	}

	team, _ := db.GetTeamByID(teamID)
	playersResults, err := db.GetTeamBattles(team.ID)
	if err != nil {
		return nil, err
	}

	var brawlerStat []model.BrawlerStat
	if err := getBrawlerStat(&pbs, &brawlerStat); err != nil {
		return nil, err
	}

	teamStat := model.TeamStat{
		Team:          team.Name,
		BattleResults: *playersResults,
		Players:       pbs,
		Brawlers:      brawlerStat,
	}

	return &teamStat, nil
}
