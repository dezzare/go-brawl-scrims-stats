package stats

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/client"
	"github.com/dezzare/go-brawl-scrims-stats/internal/service"
)

func Start(db *service.DB, c *client.Client) {
	fmt.Println("Starting Stats")

	if err := setBrawlerBase(db, c); err != nil {
		fmt.Println(err)
	}

	if err := setPlayersBase(db); err != nil {
		fmt.Println(err)
	}

	if err := setPlayersBattlelog(db, c); err != nil {
		fmt.Println(err)
	}

	showStats(db)

}

func showStats(db *service.DB) {
	players, err := db.GetPlayersFollowed()
	if err != nil {
		fmt.Println(err)
	}

	// Create map for check if
	teamMap := make(map[string]bool)
	for _, p := range *players {
		if p.TeamID != nil {
			if p.Team.Name != "" && !teamMap[p.Team.Name] {
				teamMap[p.Team.Name] = true
				showTeamStats(p.Team.ID, db)
			}
		}
	}
}
