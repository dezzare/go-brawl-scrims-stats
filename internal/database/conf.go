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

// getDBConfig get the env configurations and return the database configuration
func getDBConfig() *DBConfig {
	var dbConfig DBConfig
	fmt.Println("Getting ENV variables for DB confing")
	dbConfig.UserName = os.Getenv("DB_USERNAME")
	dbConfig.Password = os.Getenv("DB_PASSWORD")
	dbConfig.DBName = os.Getenv("DB_NAME")
	dbConfig.Host = os.Getenv("DB_HOST")
	dbConfig.Port = os.Getenv("DB_PORT")
	return &dbConfig
}

func getDsn(dbc *DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbc.Host, dbc.UserName, dbc.Password, dbc.DBName, dbc.Port)

}
