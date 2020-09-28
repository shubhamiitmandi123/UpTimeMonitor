package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig() *dbConfig {
	_, d := os.LookupEnv("DOCKER")
	godotenv.Load(".env")
	dbConfig := dbConfig{
		Host:     "localhost",
		Port:     3306,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
	}
	if d {
		dbConfig.Host = "host.docker.internal"
	}
	return &dbConfig

}

func dbURL(dbConfig *dbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
