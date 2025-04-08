package util

import "github.com/dezzare/go-brawl-scrims-stats/internal/database/registry"

func isSameTeam(a string, b string) bool {
	p1, _ := registry.GetPlayerByTag(a)
	p2, _ := registry.GetPlayerByTag(b)

	if p1.Team == p2.Team {
		return true
	}
	return false
}
