package database

import (
	"fmt"
	"log"
	"time"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/migrate"
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/repository"
	"github.com/dezzare/go-brawl-scrims-stats/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New return the DB service
// TODO: connect to postgres OR mysql (setted in env var)
func New() (*service.DB, error) {
	dbConfig := getDBConfig()
	dsn := getDsn(dbConfig)

	db, err := open(dsn)
	if err != nil {
		return nil, err
	}

	migrate.AutoMigrate(db)
	postgres, err := repository.NewPostgresRepository(db)
	if err != nil {
		return nil, err
	}

	return service.NewDB(postgres), nil
}

// Inject is used for dependency injection to gin context
func Inject(db *service.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

// Open try connect to a Postgres database
func open(dsn string) (*gorm.DB, error) {
	fmt.Println("OPENNING DB CONNECTION")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil || db == nil {
		log.Println("DB: waiting for the database to become available.")

		for i := 1; i <= 10; i++ {
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

			if db != nil && err == nil {
				return db, nil
			}

			time.Sleep(5 * time.Second)
		}

		if err != nil || db == nil {
			fmt.Printf("Error connecting DB: %v\n", err)
			return nil, err
		}
	}

	fmt.Println("DB connected")
	return db, nil

}
