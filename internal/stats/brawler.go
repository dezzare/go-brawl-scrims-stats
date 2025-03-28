package stats

import (
	"encoding/json"
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/client"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
)

// setBrawlersBase request to API all Brawlers and save to DB
func setBrawlersBase() {
	c := client.ClientConn()

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
