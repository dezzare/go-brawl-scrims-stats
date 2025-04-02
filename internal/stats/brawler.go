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
	fmt.Println("Setting Brawler DB")
	c := client.ClientConn()

	var model entity.Brawlers
	if err := json.Unmarshal(c.GetBrawlers(), &model); err != nil {
		fmt.Println("Error unmarshal brawler: ", err)
		return
	}

	db := database.Db()
	for _, v := range model.Brawler {
		if err := db.Save(&v).Error; err != nil {
			fmt.Println("Error saving brawler: ", v, err)
		}
	}
	fmt.Println("Brawler DB ready")
}
