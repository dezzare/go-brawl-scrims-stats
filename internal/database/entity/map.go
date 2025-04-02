package entity

import "gorm.io/gorm"

type Map struct {
	gorm.Model
	Name   string
	ModeID uint
	Mode   Mode
}
