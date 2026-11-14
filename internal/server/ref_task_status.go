package server

import (
	ref_task_status2 "to-do-microservices/internal/domain/ref_task_status"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RefTaskStatusRoutes(echo *echo.Echo, db *gorm.DB) {
	refTaskStatusRepository := ref_task_status2.NewRefTaskStatusRepository(db)
	refTaskStatusHandler := ref_task_status2.NewRefTaskStatusHandler(refTaskStatusRepository)

	routes := echo.Group("/ref-task-status")
	routes.GET("/list", refTaskStatusHandler.GetListRefTaskStatus)
	routes.POST("/create", refTaskStatusHandler.CreateRefTaskStatus)
}
