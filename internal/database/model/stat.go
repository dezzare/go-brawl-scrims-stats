package model

type TeamStat struct {
	Team          string
	BattleResults []PlayerResult
	Players       []PlayerBrawlerStat
	Brawlers      []BrawlerStat
}

type PlayerBrawlerStat struct {
	PlayerName string
	Brawlers   []BrawlerStat
}

type BrawlerStat struct {
	Name      string
	Victories uint
	Defeat    uint
	Draw      uint
}
