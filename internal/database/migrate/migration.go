package migrate

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
	"gorm.io/gorm"
)

// populateRegistry add all Migration Models to one map
func populateRegistry() map[string]interface{} {
	registry := make(map[string]interface{})

	registry["Ping"] = entity.Ping{}

	return registry
}

// AutoMigrate Migrate all Schemas
func AutoMigrate(db *gorm.DB) error {
	fmt.Println("Populating")
	r := populateRegistry()
	for _, v := range r {
		err := db.AutoMigrate(v)
		fmt.Printf("%v", v)
		if err != nil {
			return err
		}
	}
	fmt.Println("Migration completed")
	return nil
}
