package routes

import (
	"to-do-microservices/handlers/ref_task_status"
	"to-do-microservices/repositories"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RefTaskStatusRoutes(echo *echo.Echo, db *gorm.DB) {
	refTaskStatusRepository := repositories.NewRefTaskStatusRepository(db)
	refTaskStatusHandler := ref_task_status.NewRefTaskStatusHandler(refTaskStatusRepository)

	routes := echo.Group("/ref-task-status")
	routes.GET("/list", refTaskStatusHandler.GetListRefTaskStatus)
}
