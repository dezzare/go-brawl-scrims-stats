package database

import (
	"os"
)

type DBConfig struct {
	UserName string
	Password string
	DBName   string
	Port     string
	Host     string
}

// dbConfig is the global DB configuration
var dbConfig DBConfig

// setDBConfig get the env configurations and return the database configuration
func setDBConfig() {

	dbConfig.UserName = os.Getenv("UNAMEDB")
	dbConfig.Password = os.Getenv("PASSDB")
	dbConfig.DBName = os.Getenv("DBNAME")
	dbConfig.Port = os.Getenv("DBPORT")
	dbConfig.Host = os.Getenv("HOSTDB")


}
