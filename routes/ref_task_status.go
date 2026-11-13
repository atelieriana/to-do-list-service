package routes

import (
	"to-do-microservices/handlers/ref_task_status"

	"github.com/labstack/echo/v5"
)

func RefTaskStatusRoutes(echo *echo.Echo) {
	routes := echo.Group("/ref-task-status")
	routes.GET("/list", ref_task_status.GetListRefTaskStatus)
}
