package server

import "github.com/dezzare/go-brawl-scrims-stats/internal/router"

func Start() {
	router.Run(":5000")
}
