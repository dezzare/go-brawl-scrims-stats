package migrate

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"gorm.io/gorm"
)

// getEntityModels return slice of all Migration Models
func getEntityModels() []interface{} {
	return []interface{}{
		&entity.Ping{}, // Just for testing
		&entity.Brawler{},
		&entity.Event{},
		&entity.Map{},
		&entity.Mode{},
		&entity.Player{},
		&entity.Team{},
	}
}

// AutoMigrate Migrate all Schemas
func AutoMigrate(db *gorm.DB) error {
	m := getEntityModels()
	for _, v := range m {
		err := db.AutoMigrate(v)
		if err != nil {
			return err
		}
	}
	return nil
}
