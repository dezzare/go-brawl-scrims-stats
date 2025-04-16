package model

import "gorm.io/gorm"

type Map struct {
	gorm.Model
	Name   string `gorm:"index:idx_mode_map,unique"`
	ModeID uint   `gorm:"index:idx_mode_map,unique"`
	Mode   Mode   `gorm:"foreignKey:ModeID"`
}
