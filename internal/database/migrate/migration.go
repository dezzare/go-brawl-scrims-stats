package migrate

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
	"gorm.io/gorm"
)

// getEntityModels return slice of all Migration Models
func getEntityModels() []interface{} {
	return []interface{}{
		&model.Ping{}, // Just for testing
		&model.Battle{},
		&model.Brawler{},
		&model.Map{},
		&model.Mode{},
		&model.Player{},
		&model.PlayerResult{},
		&model.Team{},
	}
}

// AutoMigrate Migrate all Schemas
func AutoMigrate(db *gorm.DB) error {
	fmt.Println("Starting Migration")
	m := getEntityModels()
	for _, v := range m {
		err := db.AutoMigrate(v)
		if err != nil {
			fmt.Println("Migration error: ", err)
			return err
		}
	}

	fmt.Println("Migration Completed")
	return nil
}
