package entity

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name    string   `json:"name"`
	Tag     string   `json:"tag"`
	Battles []Battle `gorm:"many2many:player_battles;"`
	TeamID  uint
}
