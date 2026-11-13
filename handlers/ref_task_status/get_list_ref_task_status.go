package ref_task_status

import (
	"fmt"
	"net/http"
	"to-do-microservices/constant"
	"to-do-microservices/models"
	"to-do-microservices/repositories"

	"github.com/labstack/echo/v5"
)

func GetListRefTaskStatus(echoCtx *echo.Context) error {
	refTaskStatus, err := repositories.GetListTaskStatus()
	if err != nil {
		fmt.Println(err)
	}

	if len(refTaskStatus) == 0 {
		return echoCtx.JSON(http.StatusNotFound, constant.EmptyResponse{
			HttpCode: http.StatusNotFound,
			Message:  "Ref Status Not Found",
		})
	}

	return echoCtx.JSON(http.StatusOK, constant.ArrayResponse[models.RefTaskStatus]{
		HttpCode: http.StatusOK,
		Message:  "Success",
		Data:     refTaskStatus,
	})
}
