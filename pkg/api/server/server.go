package server

import (
	"os"

	"github.com/dezzare/go-brawl-scrims-stats/internal/router"
)

func Start() {
	router.Run(":" + os.Getenv("APP_PORT"))
}
