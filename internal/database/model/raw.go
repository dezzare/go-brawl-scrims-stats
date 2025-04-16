package model

type RawBrawler struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Power    int    `json:"power"`
	Trophies int    `json:"trophies"`
}
type RawBrawlers struct {
	Brawler []RawBrawler `json:"items"`
}

type RawPlayer struct {
	Tag     string     `json:"tag"`
	Name    string     `json:"name"`
	Brawler RawBrawler `json:"brawler"`
	Team    string     `json:"team"`
	Follow  bool
}

type RawPlayers struct {
	Player []RawPlayer `json:"items"`
}

type RawBattle struct {
	Mode       string        `json:"mode"`
	Type       string        `json:"type"`
	Result     string        `json:"result"`
	Duration   int           `json:"duration"`
	StarPlayer *RawPlayer    `json:"starPlayer"`
	Teams      [][]RawPlayer `json:"teams"`
}

type RawEvent struct {
	ID   int    `json:"id"`
	Mode string `json:"mode"`
	Map  string `json:"map"`
}

type RawMatch struct {
	BattleTime string    `json:"battleTime"`
	Event      RawEvent  `json:"event"`
	Battle     RawBattle `json:"battle"`
}

type RM struct {
	Items []RawMatch `json:"items"`
}
