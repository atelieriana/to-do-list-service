package routes

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Register(echo *echo.Echo, db *gorm.DB) {
	// Register routes below
	RefTaskStatusRoutes(echo, db)
}
