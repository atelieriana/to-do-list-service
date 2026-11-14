package main

import (
	"fmt"
	"log"
	"to-do-microservices/internal/config"
	"to-do-microservices/internal/database"
	"to-do-microservices/internal/server"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	// Init echo and .env
	runEcho := echo.New()
	port := config.GetAppPort()

	// Setup Middleware
	runEcho.Use(middleware.RequestLogger())
	runEcho.Use(middleware.Recover())

	// Open Connection to Database
	db, err := database.ConnectPostgres()
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		if errClose := sqlDB.Close(); errClose != nil {
			fmt.Println(errClose)
		}
	}()

	// Registering Route
	server.Register(runEcho, db)
	if err := runEcho.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
