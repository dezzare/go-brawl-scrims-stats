package stats

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/client"
)

func Start() {
	c := client.ClientConn()
	setBrawlersBase(c)
	// setPlayersStats(c)
	// setTeamStats(c)
}

// func setPlayersStats(c *client.Client) {
// 	tags := entity.GetAllPlayers()
// 	c.GetPlayer("str")

// }
