package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Schema   string
}

func connectDatabase() (*gorm.DB, error) {
	dsn := "host=" + getDataFromEnv().Host +
		" user=" + getDataFromEnv().Username +
		" password=" + getDataFromEnv().Password +
		" dbname=" + getDataFromEnv().Database +
		" port=" + getDataFromEnv().Port +
		" search_path=" + getDataFromEnv().Schema +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbSQL, err := db.DB()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := dbSQL.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func getDataFromEnv() Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	schema := os.Getenv("DB_SCHEMA")

	return Database{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
		Schema:   schema,
	}
}
