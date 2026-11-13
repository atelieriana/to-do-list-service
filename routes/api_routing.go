package routes

import (
	"github.com/labstack/echo/v5"
)

func Register(echo *echo.Echo) {
	// Register routes below
	RefTaskStatusRoutes(echo)
}
