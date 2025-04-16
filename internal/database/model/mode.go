package model

import "gorm.io/gorm"

type Mode struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	Maps []Map  `gorm:"foreignKey:ModeID"`
}
