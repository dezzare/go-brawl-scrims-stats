package main

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/client"
	"github.com/dezzare/go-brawl-scrims-stats/internal/conf"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database"
	"github.com/dezzare/go-brawl-scrims-stats/internal/stats"
	"github.com/dezzare/go-brawl-scrims-stats/pkg/api/server"
)

func main() {

	database.Start()
	client.Start()
	stats.Start()
	server.Start()
}

func init() {
	conf.LoadEnvConfig()
}
