package entity

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name   string
	Tag    string
	TeamID uint
	Events []Event `gorm:"many2many:player_events;"`
}
