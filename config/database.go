package config

import (
	"context"
	"log"
	"os"
	"sync"
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

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
	dbErr      error
)

func ConnectDatabase() (*gorm.DB, error) {
	conf := getDataFromEnv()
	dsn := "host=" + conf.Host +
		" user=" + conf.Username +
		" password=" + conf.Password +
		" dbname=" + conf.Database +
		" port=" + conf.Port +
		" search_path=" + conf.Schema +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbSQL, err := db.DB()
	if err != nil {
		return nil, err
	}

	dbSQL.SetMaxIdleConns(10)
	dbSQL.SetMaxOpenConns(100)
	dbSQL.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := dbSQL.PingContext(ctx); err != nil {
		_ = dbSQL.Close()
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
