package main

import (
	"log"
	"os"
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

	// Registering Route
	routes.Register(runEcho)
	if err := runEcho.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
