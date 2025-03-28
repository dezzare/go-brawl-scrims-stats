package entity

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name    string
	Players []Player
}
