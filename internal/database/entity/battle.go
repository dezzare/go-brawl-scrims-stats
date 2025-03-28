package entity

import "gorm.io/gorm"

type Battle struct {
	gorm.Model
	Mode       Mode
	Map        Map
	Result     string
	BattleTime string
	Brawlers   []BrawlerPlayed
}
