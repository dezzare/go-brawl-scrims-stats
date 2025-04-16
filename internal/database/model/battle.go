package model

import (
	"gorm.io/gorm"
)

type Battle struct {
	gorm.Model
	BattleTime string         `gorm:"index:idx_battle_unique,unique"`
	MapID      uint           `gorm:"index:idx_battle_unique,unique"`
	ModeID     uint           `gorm:"index:idx_battle_unique,unique"`
	Map        Map            `gorm:"foreignKey:MapID"`
	Mode       Mode           `gorm:"foreignKey:ModeID"`
	Results    []PlayerResult `gorm:"foreignKey:BattleID"`
}
