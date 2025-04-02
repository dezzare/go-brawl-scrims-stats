package util

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/handler"
)

func isSameTeam(a string, b string) bool {
	p1, _ := handler.GetPlayerByTag(a)
	p2, _ := handler.GetPlayerByTag(b)

	if p1.Team == p2.Team {
		return true
	}
	return false
}
