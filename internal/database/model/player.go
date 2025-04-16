package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name    string         `json:"name"`
	Tag     string         `json:"tag"`
	Battles []PlayerResult `gorm:"many2many:player_battles;"`
	Follow  bool           `gorm:"default:false"`
	Team    *Team          `gorm:"foreignKey:TeamID"`
	TeamID  *uint
}

type PlayerResult struct {
	gorm.Model
	Result    string
	PlayerID  uint    `gorm:"index;uniqueIndex:idx_player_battle"`
	BattleID  uint    `gorm:"index;uniqueIndex:idx_player_battle"`
	BrawlerID uint    `gorm:"index"`
	Player    Player  `gorm:"foreignKey:PlayerID"`
	Battle    Battle  `gorm:"foreignKey:BattleID"`
	Brawler   Brawler `gorm:"foreignKey:BrawlerID"`
}

func (p *Player) ToString() string {
	return fmt.Sprintf("Player: %s, Tag: %s, Team: %s, Follow: %v", p.Name, p.Tag, p.Team.Name, p.Follow)
}

func (p *Player) Update(pn Player) {
	p.Name = pn.Name
	p.Team = pn.Team
	p.Tag = pn.Tag
	p.Follow = pn.Follow
}
