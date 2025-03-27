package entity

import (
	"gorm.io/gorm"
)

type Brawler struct {
	gorm.Model
	ID     uint     `gorm:"primaryKey" json:"id"`
	Name   string   `json:"name"`
	Events []Battle `gorm:"many2many:brawler_events"`
}

type Brawlers struct {
	Brawler []Brawler `json:"items"`
}
