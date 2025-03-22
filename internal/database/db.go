package database

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/migrate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// dbConn is the global gorm.DB connection provider.
var dbConn DbConn

type DbConn struct {
	Dsn  string
	db   *gorm.DB
	once sync.Once
}

func (dbc *DbConn) setDsn() {
	dbc.Dsn = fmt.Sprintf(
		//	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		// dbc.Host, dbc.UserName, dbc.Password, dbc.DBName, dbc.Port)
		"host=%s user=%s password=%s port=%s sslmode=disable",
		dbConfig.Host, dbConfig.UserName, dbConfig.Password, dbConfig.Port)
}

// Open connect to a database
func (dbc *DbConn) open() error {

	db, err := gorm.Open(postgres.Open(dbc.Dsn), &gorm.Config{})
	if err != nil || db == nil {
		log.Println("DB: waiting for the database to become available.")

		for i := 1; i <= 10; i++ {
			db, err := gorm.Open(postgres.Open(dbc.Dsn), &gorm.Config{})

			if db != nil && err == nil {
				dbc.db = db
				break
			}

			time.Sleep(5 * time.Second)
		}

		if err != nil || db == nil {
			return err
		}
	}

	dbc.db = db
	return nil

}

func Start() {
	setDBConfig()
	dbConn.setDsn()

	if err := dbConn.open(); err != nil {
		log.Fatal("Failed to initialize database")
	}

	migrate.AutoMigrate(dbConn.db)

}

// Db returns the default *gorm.DB connection.
func Db() *gorm.DB {
	return dbConn.db
}
