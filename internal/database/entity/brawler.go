package entity

import (
	"gorm.io/gorm"
)

type Brawler struct {
	gorm.Model
	Name   string
	Events []Event `gorm:"many2many:brawler_events"`
}
