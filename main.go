package main

import (
	"fmt"
	"log"
	"os"
	"to-do-microservices/config"
	"to-do-microservices/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	// Init echo and .env
	runEcho := echo.New()
	errDotEnv := godotenv.Load()
	if errDotEnv != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup Middleware
	runEcho.Use(middleware.RequestLogger())
	runEcho.Use(middleware.Recover())

	// Register Port
	port := os.Getenv("APP_PORT")

	// Open Connection to DB
	db, err := config.ConnectDatabase()
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		errClose := sqlDB.Close()
		if errClose != nil {
			fmt.Println(errClose)
		}
	}()

	// Registering Route
	routes.Register(runEcho)
	if err := runEcho.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
