package entity

import "gorm.io/gorm"

type Battle struct {
	gorm.Model
	Mode     Mode      `gorm:"foreignKey:Name"`
	Map      Map       `gorm:"foreignKey:Name"`
	Brawlers []Brawler `gorm:"many2many:brawler_events"`
	Players  []Player  `gorm:"many2many:player_events"`
	Result   string
}
