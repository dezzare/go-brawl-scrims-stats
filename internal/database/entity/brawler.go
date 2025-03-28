package entity

import (
	"gorm.io/gorm"
)

type Brawlers struct {
	Brawler []Brawler
}

type Brawler struct {
	gorm.Model
	ID      uint     `gorm:"primaryKey" json:"id"`
	Name    string   `json:"name"`
	Battles []Battle `gorm:"many2many:brawler_battles"`
}

type BrawlerPlayed struct {
	gorm.Model
	Player    Player
	BrawlerID uint
	BattleID  uint
}
