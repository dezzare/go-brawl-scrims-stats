package entity

import (
	"fmt"

	"gorm.io/gorm"
)

type Players struct {
	Player []Player `json:"items"`
}

type Player struct {
	gorm.Model
	Name    string         `json:"name"`
	Tag     string         `json:"tag"`
	Battles []BattleResult `gorm:"many2many:player_battles;"`
	Team    string         `json:"team"`
	Follow  bool           `gorm:"default:false"`
}

func (p *Player) ToString() string {
	return fmt.Sprintf("Player: %s, Tag: %s, Team: %s, Follow: %v", p.Name, p.Tag, p.Team, p.Follow)
}

func (p *Player) Update(pn Player) {
	p.Name = pn.Name
	p.Team = pn.Team
	p.Tag = pn.Tag
	p.Follow = pn.Follow
}
