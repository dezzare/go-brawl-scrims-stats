package stats

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/handler"
)

func Start() {
	fmt.Println("Starting Stats")
	setBrawlersBase()
	setPlayersBattlelog()
	getStats()
}

func getStats() {
	players := handler.GetAllPlayers()

	// Create map for check if
	teamMap := make(map[string]bool)
	for _, v := range *players {
		if !teamMap[v.Team] && v.Follow {
			teamMap[v.Team] = true
			showTeamStats(v.Team)
		}

	}
}
