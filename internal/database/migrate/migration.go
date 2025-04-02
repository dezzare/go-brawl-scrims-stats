package migrate

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"gorm.io/gorm"
)

// getEntityModels return slice of all Migration Models
func getEntityModels() []interface{} {
	return []interface{}{
		&entity.Ping{}, // Just for testing
		&entity.Brawler{},
		&entity.BattleResult{},
		&entity.Battle{},
		&entity.Map{},
		&entity.Mode{},
		&entity.Player{},
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
