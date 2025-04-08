package stats

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/registry"
)

func showTeamStats(teamName string) {

	teamStat, err := getTeamStat(teamName)
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
	// }
}

func getTeamStat(teamName string) (*TeamStat, error) {
	var players []entity.Player
	if err := registry.GetTeamPlayers(&players, teamName); err != nil {
		return nil, err
	}
	var playersID []uint
	for _, p := range players {
		playersID = append(playersID, p.ID)
	}

	var pbs []PlayerBrawlerStat
	if err := setPlayersBrawlerStat(&pbs, &players); err != nil {
		return nil, err
	}

	var battleResults []entity.BattleResult
	if err := registry.GetSharedBattles(playersID, &battleResults); err != nil {
		return nil, err
	}

	var brawlerStat []BrawlerStat
	if err := getBrawlerStat(&pbs, &brawlerStat); err != nil {
		return nil, err
	}

	teamStat := TeamStat{
		Team:          teamName,
		BattleResults: battleResults,
		Players:       pbs,
		Brawlers:      brawlerStat,
	}

	return &teamStat, nil
}

func getBrawlerStatsStr(b []BrawlerStat) string {
	total := 0
	str := ""
	for _, v := range b {
		str = str + fmt.Sprintf("    - 🛡️ Brawler: %-10s | ✅ Wins: %d | ❌ Losses: %d | ⚔️ Draws: %d\n",
			v.Name, v.Victories, v.Defeat, v.Draw)
		total = total + int(v.Victories) + int(v.Defeat) + int(v.Draw)
	}
	str = str + fmt.Sprintf("TOTAL = %v\n", total)
	return str
}
