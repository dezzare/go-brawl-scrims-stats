package model

import (
	"gorm.io/gorm"
)

type Brawler struct {
	gorm.Model
	Ref     uint           `json:"id"`
	Name    string         `json:"name"`
	Battles []PlayerResult `gorm:"foreignKey:BrawlerID"`
}
