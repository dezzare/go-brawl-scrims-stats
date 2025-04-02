package entity

import (
	"gorm.io/gorm"
)

type Battle struct {
	gorm.Model
	BattleTime string
	Mode       string
	Map        string
}

type BattleResult struct {
	gorm.Model
	Result    string
	BrawlerID uint
	PlayerID  uint
	BattleID  uint `gorm:"primaryKey"`
}
