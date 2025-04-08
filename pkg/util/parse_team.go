package util

import "github.com/dezzare/go-brawl-scrims-stats/internal/database/repository"

func isSameTeam(a string, b string) bool {
	p1, _ := repository.GetPlayerByTag(a)
	p2, _ := repository.GetPlayerByTag(b)

	if p1.Team == p2.Team {
		return true
	}
	return false
}
