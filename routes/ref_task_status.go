package routes

import (
	"to-do-microservices/handlers/ref_task_status"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RefTaskStatusRoutes(echo *echo.Echo, db *gorm.DB) {
	routes := echo.Group("/ref-task-status")
	routes.GET("/list", ref_task_status.GetListRefTaskStatus(db))
}
