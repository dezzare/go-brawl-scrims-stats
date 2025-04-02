package entity

import (
	"gorm.io/gorm"
)

type Brawlers struct {
	Brawler []Brawler `json:"items"`
}

type Brawler struct {
	gorm.Model
	ID      uint           `gorm:"primaryKey" json:"id"`
	Name    string         `json:"name"`
	Battles []BattleResult `gorm:"many2many:brawler_battles;"`
}
