package stats

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/client"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"github.com/dezzare/go-brawl-scrims-stats/pkg/util"
)

func loadPlayers() {
	//TODO load from yml or toml or json
}

func requestPlayerBattlelog(playerTag string) []entity.Battle {
	c := client.ClientConn()
	r := util.ConvertToBattle(c.GetBattleLog(playerTag))
	return r
}

func setPlayerStat() {
	return
}
