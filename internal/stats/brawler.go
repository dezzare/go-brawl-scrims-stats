package stats

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/client"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
	"github.com/dezzare/go-brawl-scrims-stats/internal/service"
	"github.com/dezzare/go-brawl-scrims-stats/pkg/convert"
)

func setBrawlerBase(db *service.DB, c *client.Client) error {
	fmt.Println("[SBB] Setting Brawler DB")

	brawlers, err := convert.RawToBrawler(c.GetBrawlers())
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range *brawlers {
		if err := db.CreateBrawler(&v); err != nil {
			fmt.Println("[SBB] Error saving brawler: ", v)
			return fmt.Errorf("%w", err)
		}
	}
	fmt.Println("[SBB] Brawler DB ready")
	return nil
}

func getBrawlerStatsStr(b []model.BrawlerStat) string {
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

func getBrawlerStat(pbs *[]model.PlayerBrawlerStat, bs *[]model.BrawlerStat) error {
	pbm := make(map[string]*model.BrawlerStat)
	for _, player := range *pbs {
		for _, brawler := range player.Brawlers {
			if _, exists := pbm[brawler.Name]; !exists {
				pbm[brawler.Name] = &model.BrawlerStat{
					Name: brawler.Name,
				}
			}
			pbm[brawler.Name].Victories = pbm[brawler.Name].Victories + brawler.Victories
			pbm[brawler.Name].Draw = pbm[brawler.Name].Draw + brawler.Draw
			pbm[brawler.Name].Defeat = pbm[brawler.Name].Defeat + brawler.Defeat
		}
	}

	var brawlerStat []model.BrawlerStat
	for _, v := range pbm {
		brawlerStat = append(brawlerStat, *v)
	}
	(*bs) = brawlerStat

	return nil
}
