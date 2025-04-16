package main

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/api"
	"github.com/dezzare/go-brawl-scrims-stats/internal/client"
	"github.com/dezzare/go-brawl-scrims-stats/internal/conf"
	"github.com/dezzare/go-brawl-scrims-stats/internal/stats"
)

func main() {
	app, err := api.New()
	if err != nil {
		fmt.Println(err)
	}

	apiClient := client.New()

	stats.Start(app.DB, apiClient)
	app.Run()
}

func init() {
	conf.LoadEnvConfig()
}
