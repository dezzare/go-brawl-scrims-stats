package database

import (
	"fmt"
	"os"
)

type DBConfig struct {
	UserName string
	Password string
	DBName   string
	Host     string
	Port     string
}

// dbConfig is the global DB configuration
var dbConfig DBConfig

// setDBConfig get the env configurations and return the database configuration
func setDBConfig() {
	fmt.Println("Getting ENV variables for DB confing")
	dbConfig.UserName = os.Getenv("DB_USERNAME")
	dbConfig.Password = os.Getenv("DB_PASSWORD")
	dbConfig.DBName = os.Getenv("DB_NAME")
	dbConfig.Host = os.Getenv("DB_HOST")
	dbConfig.Port = os.Getenv("DB_PORT")

}
