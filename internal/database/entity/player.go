package entity

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name   string  `json:"name"`
	Tag    string  `josn:"tag"`
	TeamID uint    `gorm:"default:none"`
	Events []Event `gorm:"many2many:player_events;"`
}
