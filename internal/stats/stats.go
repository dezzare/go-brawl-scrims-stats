package stats

import (
	"encoding/json"
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/client"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
)

func Start() {
	c := client.ClientConn()
	setBrawlersStats(c)
	// setPlayersStats(c)
	// setTeamStats(c)
}

// setBrawlersStats request to API all Brawlers and save to DB
func setBrawlersStats(c *client.Client) {
	var model entity.Brawlers
	if err := json.Unmarshal(c.GetBrawlers(), &model); err != nil {
		fmt.Println(err)
		return
	}

	db := database.Db()
	for _, v := range model.Brawler {
		db.Save(&v)
	}

}

// func setPlayersStats(c *client.Client) {
// 	tags := entity.GetAllPlayers()
// 	c.GetPlayer("str")

// }
